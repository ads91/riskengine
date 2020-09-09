package dict

import (
	"log"
	"encoding/json"
	"io/ioutil"
	"strconv"
)

// Dict : multi-level key-value store
type Dict = map[string]interface{}

// Get : traverse a Dict in order of an array of keys
func Get(d Dict, keys []string) interface{} {
	if len(keys) == 1 {
		return d[keys[0]]
	}
	v, ok := d[keys[0]]
	if !ok {
		log.Fatal("key " + keys[0] + " not in dictionary")
	}
	return Get(v.(Dict), keys[1:])
}

// // Set : set a value in a Dict traversing it with an array of keys
// func (d Dict) Set(keys []interface{}, value interface{}, override bool) {
// 	if len(keys) == 1 {
// 		d[keys[0]] = value
// 		return
// 	}
// 	v, ok := d[keys[0]]
// 	if !ok {
// 		d[keys[0]] = Dict{}
// 	} else {
// 		log.Fatal("not allowed to override value in dict")
// 		return
// 	}
// 	v.(Dict).Set(keys[1:], value, override)
// }

// LoadFromDir : instantiate a Dict from a JSON located in a directory
func LoadFromDir(dir string) Dict {
	var d Dict
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
