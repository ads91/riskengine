package main

import (
	"log"
	"sync"

	"github.com/ads91/utils"
	"github.com/preichenberger/go-coinbasepro/Documents/dev/go/src/riskengine/pricing"
)

var env = utils.LoadFromDir("/Users/andrewsanderson/Documents/dev/go/src/riskengine/data/env.json")

func main() {
	log.Print(env)
	runLocal()
}

func runLocal() {
	var wg sync.WaitGroup
	// add some waits
	wg.Add(4)
	// price the products
	go pricing.PriceFromDir(&wg, "/Users/andrewsanderson/Documents/dev/go/src/riskengine/data/bond_01.json")
	go pricing.PriceFromDir(&wg, "/Users/andrewsanderson/Documents/dev/go/src/riskengine/data/bond_02.json")
	go pricing.PriceFromDir(&wg, "/Users/andrewsanderson/Documents/dev/go/src/riskengine/data/bond_03.json")
	go pricing.PriceFromDir(&wg, "/Users/andrewsanderson/Documents/dev/go/src/riskengine/data/europeancall_01.json")
	// wait for the above routines to return
	wg.Wait()
}
