package dict

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strconv"
)

// Dict : a hash map of interface type key to interface type value
type Dict map[interface{}]interface{}

// Get : traverse a Dict in order of an array of keys
func (d Dict) Get(keys []interface{}) interface{} {
	if len(keys) == 1 {
		return d[keys[0]]
	}
	v, ok := d[keys[0]]
	if !ok {
		log.Fatal("key " + keys[0].(string) + " not in dictionary")
	}
	return v.(Dict).Get(keys[1:])
}

// Set : set a value in a Dict traversing it with an array of keys
func (d Dict) Set(keys []interface{}, value interface{}, override bool) {
	if len(keys) == 1 {
		d[keys[0]] = value
		return
	}
	v, ok := d[keys[0]]
	if !ok {
		d[keys[0]] = Dict{}
	} else {
		log.Fatal("not allowed to override value in dict")
		return
	}
	v.(Dict).Set(keys[1:], value, override)
}

// MakeKeys : collate a number of keys into an array of keys
func MakeKeys(keys ...interface{}) []interface{} {
	return keys
}

// Dict2 (beta) : keys are string and values are interface; works better for type assertions
type Dict2 = map[string]interface{}

// Get : traverse a Dict2 in order of an array of keys
func Get(d Dict2, keys []string) interface{} {
	if len(keys) == 1 {
		return d[keys[0]]
	}
	v, ok := d[keys[0]]
	if !ok {
		log.Fatal("key " + keys[0] + " not in dictionary")
	}
	return Get(v.(Dict2), keys[1:])
}

// Set : traverse a Dict2 in order of an array of keys
func Set(d Dict2) {}

// LoadFromDir : instantiate a Dict2 from a JSON located in a directory
func LoadFromDir(dir string) Dict2 {
	var d Dict2
	// read
	file, err := ioutil.ReadFile(dir)
	// check the read
	if err != nil {
		log.Fatal("could not read file at path ", dir)
	}
	// deserialise the JSON
	json.Unmarshal([]byte(file), &d)
	// return
	return d
}

// ConvDictFloat64 : convert an interface type into a float64
func ConvDictFloat64(val interface{}) float64 {
	v, err := strconv.ParseFloat(val.(string), 64)
	if err != nil {
		log.Fatal("incorrect input type ", val)
	}
	return v
}

// ConvDictInt : convert an interface type into an int
func ConvDictInt(val interface{}) int {
	v, err := strconv.Atoi(val.(string))
	if err != nil {
		log.Fatal("incorrect input type ", val)
	}
	return v
}
