package main

import (
	"fmt"

	"github.com/nmccready/colorjson"
)

func main() {
	simpleMap := colorjson.Object{}
	simpleMap["a"] = 1
	simpleMap["b"] = "bee"
	simpleMap["c"] = []float64{1, 2, 3}
	simpleMap["d"] = []string{"one", "two", "three"}

	bytes, _ := colorjson.Marshal(simpleMap)
	fmt.Println(string(bytes))
}
