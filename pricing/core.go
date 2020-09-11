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
	var results = dict.Dict{}
	// conversion for all products
	for id, config := range trades {
		config := config.(dict.Dict)
		productType := config["type"]
		args := config["args"].(dict.Dict)
		// check product type and instantiate accordingly
		config["error"] = 0
		switch productType {
		case "bond":
			curveType := args["curve"].(string)
			curve := environment.NewCurve(curveType, env)
			bond := Bond{
				Curve:  curve,
				Coupon: dict.ConvDictFloat64(args["coupon"]),
			}
			config["price"] = bond.Price()
		case "europeancall":
			europeanCall := EuropeanCall{
				s0:  dict.ConvDictFloat64(args["startprice"]),
				K:   dict.ConvDictFloat64(args["strike"]),
				T:   dict.ConvDictInt(args["years"]),
				R:   dict.ConvDictFloat64(args["rate"]),
				Vol: dict.ConvDictFloat64(args["vol"]),
				N:   dict.ConvDictInt(args["paths"]),
			}
			config["price"] = europeanCall.Price()
		default:
			config["error"] = 1
			config["price"] = "unrecognised product type " + productType.(string)
		}
		results[id] = config
	}
	return results
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
	log.Print("recieved request to price ", d)
	Price(d, hp.Env)
}
