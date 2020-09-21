package colorjson

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

const initialDepth = 0
const valueSep = ","
const null = "null"
const startMap = "{"
const endMap = "}"
const startArray = "["
const endArray = "]"

const emptyMap = startMap + endMap
const emptyArray = startArray + endArray

var stripColorsRegEx = regexp.MustCompile(`\x1b\[[0-9;]*m`)

type Object map[string]interface{}

type Formatter struct {
	KeyColor        *color.Color
	StringColor     *color.Color
	BoolColor       *color.Color
	NumberColor     *color.Color
	NullColor       *color.Color
	StringMaxLength int
	Indent          int
	DisabledColor   bool
	HTMLEscape      bool
	RawStrings      bool
	KeyMapColors    map[string]*color.Color
}

func NewFormatter() *Formatter {
	return &Formatter{
		KeyColor:        color.New(color.FgWhite),
		StringColor:     color.New(color.FgGreen),
		BoolColor:       color.New(color.FgYellow),
		NumberColor:     color.New(color.FgCyan),
		NullColor:       color.New(color.FgMagenta),
		StringMaxLength: 0,
		DisabledColor:   false,
		Indent:          0,
		RawStrings:      false,
		KeyMapColors:    map[string]*color.Color{},
	}
}

//nolint unused
func (f *Formatter) sprintfColor(key string, c *color.Color, format string, args ...interface{}) string {
	if f.KeyMapColors[key] != nil {
		c = f.KeyMapColors[key]
	}
	if f.DisabledColor || c == nil {
		return fmt.Sprintf(format, args...)
	}
	return c.SprintfFunc()(format, args...)
}

func (f *Formatter) sprintColor(key string, c *color.Color, s string) string {
	if f.KeyMapColors[key] != nil {
		c = f.KeyMapColors[key]
	}
	if f.DisabledColor || c == nil {
		return fmt.Sprint(s)
	}
	return c.SprintFunc()(s)
}

func (f *Formatter) writeIndent(buf *bytes.Buffer, depth int) {
	buf.WriteString(strings.Repeat(" ", f.Indent*depth))
}

func (f *Formatter) writeObjSep(buf *bytes.Buffer) {
	if f.Indent != 0 {
		buf.WriteByte('\n')
	} else {
		buf.WriteByte(' ')
	}
}

func (f *Formatter) Marshal(jsonObj interface{}) ([]byte, error) {
	buffer := bytes.Buffer{}
	f.marshalValue("", jsonObj, &buffer, initialDepth)
	if f.DisabledColor {
		return stripColorsRegEx.ReplaceAll(buffer.Bytes(), []byte("")), nil
	}
	return buffer.Bytes(), nil
}

func (f *Formatter) MarshalString(jsonObj interface{}) (string, error) {
	b, err := f.Marshal(jsonObj)
	return string(b), err
}

func (f *Formatter) marshalMap(m map[string]interface{}, buf *bytes.Buffer, depth int) {
	remaining := len(m)

	if remaining == 0 {
		buf.WriteString(emptyMap)
		return
	}

	keys := make([]string, 0)
	for key := range m {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	buf.WriteString(startMap)
	f.writeObjSep(buf)

	for _, key := range keys {
		f.writeIndent(buf, depth+1)
		buf.WriteString(f.KeyColor.Sprintf("\"%s\": ", key))
		f.marshalValue(key, m[key], buf, depth+1)
		remaining--
		if remaining != 0 {
			buf.WriteString(valueSep)
		}
		f.writeObjSep(buf)
	}
	f.writeIndent(buf, depth)
	buf.WriteString(endMap)
}

func (f *Formatter) marshalArray(key string, a []interface{}, buf *bytes.Buffer, depth int) {
	if len(a) == 0 {
		buf.WriteString(emptyArray)
		return
	}

	buf.WriteString(startArray)
	f.writeObjSep(buf)

	for i, v := range a {
		f.writeIndent(buf, depth+1)
		f.marshalValue(key, v, buf, depth+1)
		if i < len(a)-1 {
			buf.WriteString(valueSep)
		}
		f.writeObjSep(buf)
	}
	f.writeIndent(buf, depth)
	buf.WriteString(endArray)
}

func (f *Formatter) marshalIntArray(key string, a []int, buf *bytes.Buffer, depth int) {
	if len(a) == 0 {
		buf.WriteString(emptyArray)
		return
	}

	buf.WriteString(startArray)
	f.writeObjSep(buf)

	for i, v := range a {
		f.writeIndent(buf, depth+1)
		f.marshalValue(key, v, buf, depth+1)
		if i < len(a)-1 {
			buf.WriteString(valueSep)
		}
		f.writeObjSep(buf)
	}
	f.writeIndent(buf, depth)
	buf.WriteString(endArray)
}

func (f *Formatter) marshalFloatArray(key string, a []float64, buf *bytes.Buffer, depth int) {
	if len(a) == 0 {
		buf.WriteString(emptyArray)
		return
	}

	buf.WriteString(startArray)
	f.writeObjSep(buf)

	for i, v := range a {
		f.writeIndent(buf, depth+1)
		f.marshalValue(key, v, buf, depth+1)
		if i < len(a)-1 {
			buf.WriteString(valueSep)
		}
		f.writeObjSep(buf)
	}
	f.writeIndent(buf, depth)
	buf.WriteString(endArray)
}

func (f *Formatter) marshalBoolArray(key string, a []bool, buf *bytes.Buffer, depth int) {
	if len(a) == 0 {
		buf.WriteString(emptyArray)
		return
	}

	buf.WriteString(startArray)
	f.writeObjSep(buf)

	for i, v := range a {
		f.writeIndent(buf, depth+1)
		f.marshalValue(key, v, buf, depth+1)
		if i < len(a)-1 {
			buf.WriteString(valueSep)
		}
		f.writeObjSep(buf)
	}
	f.writeIndent(buf, depth)
	buf.WriteString(endArray)
}

func (f *Formatter) marshalStringArray(key string, a []string, buf *bytes.Buffer, depth int) {
	if len(a) == 0 {
		buf.WriteString(emptyArray)
		return
	}

	buf.WriteString(startArray)
	f.writeObjSep(buf)

	for i, v := range a {
		f.writeIndent(buf, depth+1)
		f.marshalValue(key, v, buf, depth+1)
		if i < len(a)-1 {
			buf.WriteString(valueSep)
		}
		f.writeObjSep(buf)
	}
	f.writeIndent(buf, depth)
	buf.WriteString(endArray)
}

func (f *Formatter) marshalValue(key string, val interface{}, buf *bytes.Buffer, depth int) {
	switch v := val.(type) {
	case Object:
		f.marshalMap(v, buf, depth)
	case map[string]interface{}:
		f.marshalMap(v, buf, depth)
	case []interface{}:
		f.marshalArray(key, v, buf, depth)
	case []int:
		f.marshalIntArray(key, v, buf, depth)
	case []float64:
		f.marshalFloatArray(key, v, buf, depth)
	case []string:
		f.marshalStringArray(key, v, buf, depth)
	case []bool:
		f.marshalBoolArray(key, v, buf, depth)
	case string:
		f.marshalString(key, v, buf)
	case error:
		f.marshalString(key, v.Error(), buf)
	case float64:
		buf.WriteString(f.sprintColor(key, f.NumberColor, strconv.FormatFloat(v, 'f', -1, 64)))
	case bool:
		buf.WriteString(f.sprintColor(key, f.BoolColor, (strconv.FormatBool(v))))
	case nil:
		buf.WriteString(f.sprintColor(key, f.NullColor, null))
	case json.Number:
		buf.WriteString(f.sprintColor(key, f.NumberColor, v.String()))
	case int:
		buf.WriteString(f.sprintColor(key, f.NumberColor, strconv.Itoa(v)))
	default:
		//nolint lll
		fmt.Printf("colorjson error: unknown type of " + reflect.TypeOf(v).String() + ". If this an object cast as Object or []interface{}\n")
	}
}

func (f *Formatter) marshalString(key string, str string, buf *bytes.Buffer) {
	if !f.RawStrings {
		b := &bytes.Buffer{}

		encoder := json.NewEncoder(b)
		encoder.SetEscapeHTML(f.HTMLEscape)
		err := encoder.Encode(interface{}(str))
		if err != nil {
			str = "colorjson: error encoding string"
		} else {
			str = strings.Replace(b.String(), "\n", "", 1)
		}
	}

	if f.StringMaxLength != 0 && len(str) >= f.StringMaxLength {
		str = fmt.Sprintf("%s...", str[0:f.StringMaxLength])
	}

	// buf.WriteString(str)
	buf.WriteString(f.sprintColor(key, f.StringColor, str))
}

// Marshal JSON data with default options
func Marshal(jsonObj interface{}) ([]byte, error) {
	return NewFormatter().Marshal(jsonObj)
}
