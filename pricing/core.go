package pricing

import (
	"encoding/json"
	"http"
	"log"
	"riskengine/environment"
	"sync"

	"github.com/ads91/utils"
)

// Price : price a trade (of acceptable type)
func Price(trade utils.Dict2, env utils.Dict2) {
	var price float64
	// conversion for all products
	for id, config := range trade {
		config := config.(utils.Dict2)
		productType := config["type"]
		args := config["args"].(utils.Dict2)
		// check product type and instantiate accordingly
		switch productType {
		case "bond":
			curveType := args["curve"].(string)
			curve := environment.NewCurve(curveType, env)
			bond := Bond{
				Curve:  curve,
				Coupon: utils.ConvDictFloat64(args["coupon"]),
			}
			price = bond.Price()
		case "europeancall":
			europeanCall := EuropeanCall{
				s0:  utils.ConvDictFloat64(args["startprice"]),
				K:   utils.ConvDictFloat64(args["strike"]),
				T:   utils.ConvDictInt(args["years"]),
				R:   utils.ConvDictFloat64(args["rate"]),
				Vol: utils.ConvDictFloat64(args["vol"]),
				N:   utils.ConvDictInt(args["paths"]),
			}
			price = europeanCall.Price()
		default:
			log.Fatal("unrecognised product type: ", productType)
			return
		}
		log.Printf("price for %s of type %s = %g", id, productType, price)
	}
}

// PriceFromDir : price a trade (JSON) located in a directory
func PriceFromDir(wg *sync.WaitGroup, dir string, env utils.Dict2) {
	// don't return until method's complete
	defer wg.Done()
	Price(utils.LoadFromDir(dir), env)
}

// PriceFromHTTPRequests : price a trade (JSON) sent through an HTTP request
func PriceFromHTTPRequests(hp HTTPPricer, port string, uri string) {
	// set-up the handler
	http.HandleFunc(uri, hp.httpPricingHandler)
	// listen indefinitely
	log.Fatal(http.ListenAndServe(port, nil))
}

func (hp HTTPPricer) httpPricingHandler(w http.ResponseWriter, r *http.Request) {
	var d utils.Dict2
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
