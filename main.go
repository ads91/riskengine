package main

import (
	"log"
	"os"
	"riskengine/pricing"
	"riskengine/utils/dict"
	"sync"
)

// DEFAULTPORT : default port to listen on
var DEFAULTPORT = ":8080"

func main() {
	// register the PID
	log.Printf("process ID is %d", os.Getpid())
	// get working directory
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	// create an instance of a market data environment
	env := dict.LoadFromDir(wd + "/data/env.json")
	//runLocal(wd, env)
	runHTTP("/price", env)
}

func runLocal(wd string, env dict.Dict2) {
	var wg sync.WaitGroup
	// add some waits
	wg.Add(4)
	// price the products
	go pricing.PriceFromDir(&wg, wd+"/data/bond_01.json", env)
	go pricing.PriceFromDir(&wg, wd+"/data/bond_02.json", env)
	go pricing.PriceFromDir(&wg, wd+"/data/bond_03.json", env)
	go pricing.PriceFromDir(&wg, wd+"/data/europeancall_01.json", env)
	// wait for the above routines to return
	wg.Wait()
}

func getPort() string {
	// test for port
	port := os.Getenv("PORT")
	// default port
	if port == "" {
		port = DEFAULTPORT
	}
	log.Printf("listening on port %s", port)

	return port
}

func runHTTP(uri string, env dict.Dict2) {
	pricer := pricing.HTTPPricer{Env: env}
	// handle pricing requests
	pricing.PriceFromHTTPRequests(pricer, getPort(), uri)
}
