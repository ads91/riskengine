package pricing

import (
	"encoding/json"
	"net/http"
	"riskengine/environment"
	"riskengine/utils/dict"
	"riskengine/utils/logging"
	"sync"
)

var log = logging.GetLogger()

// Price : price a trade (of acceptable type)
func Price(trade dict.Dict2, env dict.Dict2) {
	var price float64
	// conversion for all products
	for id, config := range trade {
		config := config.(dict.Dict2)
		productType := config["type"]
		args := config["args"].(dict.Dict2)
		// check product type and instantiate accordingly
		switch productType {
		case "bond":
			curveType := args["curve"].(string)
			curve := environment.NewCurve(curveType, env)
			bond := Bond{
				Curve:  curve,
				Coupon: dict.ConvDictFloat64(args["coupon"]),
			}
			price = bond.Price()
		case "europeancall":
			europeanCall := EuropeanCall{
				s0:  dict.ConvDictFloat64(args["startprice"]),
				K:   dict.ConvDictFloat64(args["strike"]),
				T:   dict.ConvDictInt(args["years"]),
				R:   dict.ConvDictFloat64(args["rate"]),
				Vol: dict.ConvDictFloat64(args["vol"]),
				N:   dict.ConvDictInt(args["paths"]),
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
func PriceFromDir(wg *sync.WaitGroup, dir string, env dict.Dict2) {
	// don't return until method's complete
	defer wg.Done()
	Price(dict.LoadFromDir(dir), env)
}

// PriceFromHTTPRequests : price a trade (JSON) sent through an HTTP request
func PriceFromHTTPRequests(hp HTTPPricer, port string, uri string) {
	// set-up the handler
	http.HandleFunc(uri, hp.httpPricingHandler)
	// listen indefinitely
	log.Fatal(http.ListenAndServe(port, nil))
}

func (hp HTTPPricer) httpPricingHandler(w http.ResponseWriter, r *http.Request) {
	var d dict.Dict2
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
