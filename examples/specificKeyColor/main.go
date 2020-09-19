package main

import (
	"encoding/json"
	"fmt"

	"github.com/fatih/color"
	"github.com/nmccready/colorjson"
)

func main() {
	str := `{
      "unique": "Oh HI!",
      "str": "foo",
      "num": 100,
      "bool": false,
      "null": null,
      "array": ["foo", "bar", "baz"],
      "obj": { "a": 1, "b": 2 }
    }`

	var obj map[string]interface{}
	json.Unmarshal([]byte(str), &obj)

	// Make a custom formatter with indent set
	f := colorjson.NewFormatter()
	f.Indent = 4
	f.KeyMapColors["unique"] = color.New(color.FgHiMagenta)

	// Marshall the Colorized JSON
	s, _ := f.Marshal(obj)
	fmt.Println(string(s))
}
