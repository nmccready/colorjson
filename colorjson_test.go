package colorjson_test

import (
	"testing"

	faith "github.com/fatih/color"
	"github.com/hokaccha/go-prettyjson"
	"github.com/nmccready/colorjson"
)

var customFormater *colorjson.Formatter = makeCustom()

func makeCustom() *colorjson.Formatter {
	formatter := colorjson.NewFormatter()
	formatter.Indent = 2
	formatter.KeyMapColors["custom"] = faith.New(faith.FgMagenta)
	return formatter
}

func BenchmarkMarshall(b *testing.B) {
	simpleMap := make(map[string]interface{})
	simpleMap["a"] = 1
	simpleMap["b"] = "bee"
	simpleMap["c"] = [3]float64{1, 2, 3}
	simpleMap["d"] = [3]string{"one", "two", "three"}

	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		//nolint errcheck
		colorjson.Marshal(simpleMap)
	}
}

func BenchmarkPrettyJSON(b *testing.B) {
	simpleMap := make(map[string]interface{})
	simpleMap["a"] = 1
	simpleMap["b"] = "bee"
	simpleMap["c"] = [3]float64{1, 2, 3}
	simpleMap["d"] = [3]string{"one", "two", "three"}

	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		//nolint errcheck
		prettyjson.Marshal(simpleMap)
	}
}

func BenchmarkPrettyCustomKeyJSON(b *testing.B) {
	simpleMap := make(map[string]interface{})
	simpleMap["a"] = 1
	simpleMap["b"] = "bee"
	simpleMap["c"] = [3]float64{1, 2, 3}
	simpleMap["d"] = [3]string{"one", "two", "three"}
	simpleMap["custom"] = "custom"

	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		//nolint errcheck
		customFormater.Marshal(simpleMap)
	}
}
