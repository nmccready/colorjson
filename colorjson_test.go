package colorjson_test

import (
	"testing"

	faith "github.com/fatih/color"
	"github.com/hokaccha/go-prettyjson"
	"github.com/nmccready/colorjson"
	"github.com/stretchr/testify/assert"
)

var customFormater *colorjson.Formatter = makeCustom()

func makeCustom() *colorjson.Formatter {
	formatter := colorjson.NewFormatter()
	formatter.Indent = 2
	formatter.KeyMapColors["custom"] = faith.New(faith.FgMagenta)
	return formatter
}

func TestSanity(t *testing.T) {
	simpleMap := colorjson.Object{}
	simpleMap["a"] = 1
	simpleMap["b"] = "bee"
	simpleMap["c"] = []float64{1, 2, 3}
	simpleMap["d"] = []string{"one", "two", "three"}

	f := colorjson.NewFormatter()
	f.DisabledColor = true

	str, _ := f.MarshalString(simpleMap)
	assert.Equal(t, `{ "a": 1, "b": "bee", "c": [ 1, 2, 3 ], "d": [ "one", "two", "three" ] }`, str)
}

func BenchmarkMarshall(b *testing.B) {
	simpleMap := make(map[string]interface{})
	simpleMap["a"] = 1
	simpleMap["b"] = "bee"
	simpleMap["c"] = []float64{1, 2, 3}
	simpleMap["d"] = []string{"one", "two", "three"}

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
	simpleMap["c"] = []float64{1, 2, 3}
	simpleMap["d"] = []string{"one", "two", "three"}

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
	simpleMap["c"] = []float64{1, 2, 3}
	simpleMap["d"] = []string{"one", "two", "three"}
	simpleMap["custom"] = "custom"

	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		//nolint errcheck
		customFormater.Marshal(simpleMap)
	}
}
