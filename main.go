package main

import (
	"log"

	"github.com/ads91/utils"
)

var env = utils.LoadFromDir("/Users/andrewsanderson/Documents/dev/go/src/riskengine/data/env.json")

func main() {
	log.Print(env)
}
