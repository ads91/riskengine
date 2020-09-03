package main

import (
	"log"
	"os"
	"riskengine/pricing"
	"riskengine/utils/dict"
	"sync"
)

var env = dict.LoadFromDir("/Users/andrewsanderson/Documents/dev/go/src/riskengine/data/env.json")

func main() {
	// register the PID
	log.Printf("process ID is %d", os.Getpid())
	runLocal()
	//runHTTP(":8080", "/price")
}

func runLocal() {
	var wg sync.WaitGroup
	// add some waits
	wg.Add(4)
	// price the products
	go pricing.PriceFromDir(&wg, "/Users/andrewsanderson/Documents/dev/go/src/riskengine/data/bond_01.json", env)
	go pricing.PriceFromDir(&wg, "/Users/andrewsanderson/Documents/dev/go/src/riskengine/data/bond_02.json", env)
	go pricing.PriceFromDir(&wg, "/Users/andrewsanderson/Documents/dev/go/src/riskengine/data/bond_03.json", env)
	go pricing.PriceFromDir(&wg, "/Users/andrewsanderson/Documents/dev/go/src/riskengine/data/europeancall_01.json", env)
	// wait for the above routines to return
	wg.Wait()
}

func runHTTP(port string, uri string) {
	pricer := pricing.HTTPPricer{Env: env}
	pricing.PriceFromHTTPRequests(pricer, port, uri)
}
