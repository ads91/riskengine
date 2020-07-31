package main

import (
	"dict"
	"riskengine/pricing"
	"sync"
)

var env = dict.LoadFromDir("/Users/andrewsanderson/Documents/dev/go/src/riskengine/data/env.json")

func main() {
	runLocal()
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
