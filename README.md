# ColorJSON: [![build status][travis-image]][travis-url]

## The Fast Color JSON Marshaller for Go

![ColorJSON Output](https://i.imgur.com/pLtCXhb.png)
What is this?

---

Clone from [original](http://github.com/TylerBrock/colorjson).

This package is based heavily on hokaccha/go-prettyjson but has some noticible differences:

- Over twice as fast (recursive descent serialization uses buffer instead of string concatenation)
  ```
  BenchmarkColorJSONMarshall-4     500000      2498 ns/op
  BenchmarkPrettyJSON-4            200000      6145 ns/op
  ```
- more customizable (ability to have zero indent, print raw json strings, etc...)
- better defaults (less bold colors)

ColorJSON was built in order to produce fast beautiful colorized JSON output for [Saw](http://github.com/TylerBrock/saw).

## Installation

```sh
go get -u github.com/nmccready/colorjson
```

## Usage

Setup

```go
import "github.com/nmccready/colorjson"

str := `{
  "str": "foo",
  "num": 100,
  "bool": false,
  "null": null,
  "array": ["foo", "bar", "baz"],
  "obj": { "a": 1, "b": 2 },
  "unique": "Oh HI!"
}`

// Create an intersting JSON object to marshal in a pretty format
var obj map[string]interface{}
json.Unmarshal([]byte(str), &obj)
```

Vanilla Usage

```go
s, _ := colorjson.Marshal(obj)
fmt.Println(string(s))
```

Customization (Custom Indent)

```go
f := colorjson.NewFormatter()
f.Indent = 2
f.KeyMapColors["unique"] = color.New(color.FgHiMagenta)

s, _ := f.Marshal(v)
fmt.Println(string(s))
```

[travis-image]: https://img.shields.io/travis/nmccready/colorjson.svg
[travis-url]: https://travis-ci.org/nmccready/colorjson
