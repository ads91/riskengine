package pricing

import (
	"encoding/json"
	"log"
	"net/http"
	"riskengine/environment"
	"riskengine/utils/dict"
	"sync"
)

// Price : price a dictionary of trades (of acceptable types)
func Price(trades dict.Dict, env dict.Dict) dict.Dict {
	var wg sync.WaitGroup
	var results = dict.Dict{}
	// create a channel and add some waits
	ch := make(chan dict.Dict, len(trades))
	wg.Add(len(trades))
	// loop the trades to price them
	for id, trade := range trades {
		go price(&wg, id, trade.(dict.Dict), env, ch)
	}
	// wait for the trades to price
	wg.Wait()
	// loop the channel until no more results to retrieve
	for i := 0; i < len(trades); i++ {
		result := <-ch
		results[result["id"].(string)] = result["trade"]
	}
	return results
}

func price(wg *sync.WaitGroup, id string, trade dict.Dict, env dict.Dict, ch chan dict.Dict) {
	var result = dict.Dict{}
	defer wg.Done()
	log.Print("recieved request to price ", trade)
	productType := trade["type"]
	args := trade["args"].(dict.Dict)
	// check product type and instantiate accordingly
	trade["error"] = false
	switch productType {
	case "bond":
		curveType := args["curve"].(string)
		curve := environment.NewCurve(curveType, env)
		bond := Bond{
			Curve:  curve,
			Coupon: dict.ConvDictFloat64(args["coupon"]),
		}
		trade["price"] = bond.Price()
	case "europeancall":
		europeanCall := EuropeanCall{
			s0:  dict.ConvDictFloat64(args["startprice"]),
			K:   dict.ConvDictFloat64(args["strike"]),
			T:   dict.ConvDictInt(args["years"]),
			R:   dict.ConvDictFloat64(args["rate"]),
			Vol: dict.ConvDictFloat64(args["vol"]),
			N:   dict.ConvDictInt(args["paths"]),
		}
		trade["price"] = europeanCall.Price()
	default:
		trade["error"] = true
		trade["price"] = "unrecognised product type " + productType.(string)
	}
	result["id"] = id
	result["trade"] = trade
	ch <- result
}

// PriceFromDir : price some trades (JSON) located in a directory
func PriceFromDir(wg *sync.WaitGroup, dir string, env dict.Dict) {
	// don't return until method's complete
	defer wg.Done()
	log.Print(Price(dict.LoadFromDir(dir), env))
}

// PriceFromHTTPRequests : price some trades (JSON) sent through an HTTP request
func PriceFromHTTPRequests(hp HTTPPricer, port string, uri string) {
	// set-up the handler
	http.HandleFunc(uri, hp.httpPricingHandler)
	// listen indefinitely
	log.Fatal(http.ListenAndServe(port, nil))
}

func (hp HTTPPricer) httpPricingHandler(w http.ResponseWriter, r *http.Request) {
	var d dict.Dict
	// parse form
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	// deserialise the JSON into a struct
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&d)
	if err != nil {
		log.Fatal(err)
	}
	// call the pricer
	//log.Print("recieved request to price ", d)
	results := Price(d, hp.Env)
	// serialise the pricing results
	js, err := json.Marshal(results)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// return the JSON
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
