package types

import (
	"sort"
	"strings"
)

type JsonObject struct {
	Value map[string]*JsonElement
}

func NewJsonObject() *JsonObject {
	return &JsonObject{ Value: make(map[string]*JsonElement) }
}

func (o *JsonObject) Serialize(indent int) string {
	var objString strings.Builder
	objString.WriteString("{ ")
	if indent > 0 {
		objString.WriteString("\n")
	}

	tab := ""
	if indent > 0 {
		tab = strings.Repeat("\t", indent)
	}

	_sz := len(o.Value)
	keys := o.GetKeys()
	sort.Strings(keys)

	for _,k := range keys {
		objString.WriteString(tab + "\"" + k + "\": ")
		if indent > 0 {
			objString.WriteString((*o.Get(k)).Serialize(indent + 1))
		} else {
			objString.WriteString((*o.Get(k)).Serialize(indent))
		}

		_sz -=1
		if _sz != 0 {
			objString.WriteString(", ")
		}

		if indent > 0 {
			objString.WriteString("\n")
		}
	}

	if indent > 0 {
		objString.WriteString(strings.Repeat("\t", indent - 1) + "}")
	} else {
		objString.WriteString(" }")
	}

	return objString.String()
}

func (o *JsonObject) GetSize() int {
	return len(o.Value)
}

func (o *JsonObject) Contains(key string) bool {
	return o.Value[key] != nil
}

func (o *JsonObject) Get(key string) *JsonElement {
	if o.Contains(key) {
		return o.Value[key]
	} else {
		return nil
	}
}

func (o *JsonObject) Set(key string, value JsonElement) {
	o.Value[key] = &value
}

func (o *JsonObject) GetKeys() []string {
	keys := make([]string, 0, o.GetSize())
	for k := range o.Value {
		keys = append(keys, k)
	}
	return keys
}

func (o *JsonObject) GetArrayValue() *JsonArray {
	return nil
}

func (o *JsonObject) GetBoolValue() *JsonBool {
	return nil
}

func (o *JsonObject) GetFloatValue() *JsonFloat {
	return nil
}

func (o *JsonObject) GetIntValue() *JsonInt {
	return nil
}

func (o *JsonObject) GetNullValue() *JsonNull {
	return nil
}

func (o *JsonObject) GetObjectValue() *JsonObject {
	return o
}

func (o *JsonObject) GetStringValue() *JsonString {
	return nil
}
