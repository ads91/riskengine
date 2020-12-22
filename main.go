package main

import (
	"log"
	"os"
	"riskengine/config"
	"riskengine/pricing"
	"riskengine/utils/dict"
	"runtime"
	"sync"
)

func main() {
	// print the number of procs
	log.Print("number of cores = ", runtime.NumCPU())
	// register the PID
	log.Printf("process ID is %d", os.Getpid())
	// get working directory
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	// set working directory as env variable
	os.Setenv(config.WORKING_DIR, wd)
	// create an instance of a market data environment
	env := dict.LoadFromDir(wd + "/data/env.json")
	//runLocal(env)
	runHTTP("/price", env)
}

func runLocal(env dict.Dict) {
	var wg sync.WaitGroup
	// add some waits
	wg.Add(1)
	// price the products
	wd := os.Getenv(config.WORKING_DIR)
	go pricing.PriceFromDir(&wg, wd+"/data/trades.json", env)
	// wait for the above routines to return
	wg.Wait()
}

func getPort() string {
	// test for port
	port := os.Getenv("PORT")
	// default port
	if port == "" {
		port = config.DEFAULT_PORT
	}
	log.Printf("listening on port %s", port)
	return ":" + port
}

func runHTTP(uri string, env dict.Dict) {
	pricer := pricing.HTTPPricer{Env: env}
	// handle pricing requests
	pricing.PriceFromHTTPRequests(pricer, getPort(), uri)
}
