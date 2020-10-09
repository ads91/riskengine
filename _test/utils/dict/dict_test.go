package dict_test

import (
	"riskengine/utils/dict"
	"testing"
)

func TestDict(t *testing.T) {
	d := dict.Dict{}
	d["a"] = 123
	if d["a"].(int) != 123 {
		t.Fail()
	}
}

func TestGet(t *testing.T) {
	d := dict.Dict{}
	// trivial case
	d["a"] = 123
	if dict.Get(d, []string{"a"}) != 123 {
		t.Fail()
	}
	// multi-level case
	d["b"] = dict.Dict{}
	d2 := d["b"].(dict.Dict)
	d2["a"] = 123
	if dict.Get(d, []string{"b", "a"}) != 123 {
		t.Fail()
	}
	// intentional error
	dict.Get(d, []string{"c"})
}

// // LoadFromDir : instantiate a Dict from a JSON located in a directory
// func LoadFromDir(dir string) Dict {
// 	var d Dict
// 	// read
// 	file, err := ioutil.ReadFile(dir)
// 	// check the read
// 	if err != nil {
// 		log.Fatal("could not read file at path ", dir)
// 	}
// 	// deserialise the JSON
// 	json.Unmarshal([]byte(file), &d)
// 	// return
// 	return d
// }

// // ConvDictFloat64 : convert an interface type into a float64
// func ConvDictFloat64(val interface{}) float64 {
// 	v, err := strconv.ParseFloat(val.(string), 64)
// 	if err != nil {
// 		log.Fatal("incorrect input type ", val)
// 	}
// 	return v
// }

// // ConvDictInt : convert an interface type into an int
// func ConvDictInt(val interface{}) int {
// 	v, err := strconv.Atoi(val.(string))
// 	if err != nil {
// 		log.Fatal("incorrect input type ", val)
// 	}
// 	return v
// }
