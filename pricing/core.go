package pricing

import (
	"log"
	"riskengine/environment"

	"github.com/ads91/utils"
)

// exposed
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
				S0:  utils.ConvDictFloat64(args["startprice"]),
				K:   utils.ConvDictFloat64(args["strike"]),
				T:   utils.ConvDictInt(args["years"]),
				R:   utils.ConvDictFloat64(args["rate"]),
				Vol: utils.ConvDictFloat64(args["vol"]),
				N:   utils.ConvDictInt(args["paths"]),
			}
			price = europeanCall.Price()
		default:
			log.Fatal("unrecognised product type: ", productType)
		}
		log.Printf("price for %s of type %s = %g", id, productType, price)
	}
}
