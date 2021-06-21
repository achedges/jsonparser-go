package types

import "strings"

type JsonArray struct {
	Value []JsonElement
}

func NewJsonArray() *JsonArray {
	return &JsonArray{	Value: make([]JsonElement, 0, 10) }
}

func (a *JsonArray) Serialize(indent int) string {
	var jsonArray strings.Builder
	jsonArray.WriteString("[ ")

	if indent > 0 {
		jsonArray.WriteString("\n")
	}

	tab := ""
	if indent > 0 {
		tab = strings.Repeat("\t", indent)
	}

	_sz := len(a.Value)
	for i, e := range a.Value {
		jsonArray.WriteString(tab + e.Serialize(indent))
		if i != _sz - 1 {
			jsonArray.WriteString(", ")
		}
		if indent > 0 {
			jsonArray.WriteString("\n")
		}
	}

	if indent > 0 {
		jsonArray.WriteString(tab[0:indent])
	} else {
		jsonArray.WriteString(" ")
	}

	jsonArray.WriteString("]")

	return jsonArray.String()
}

func (a *JsonArray) GetSize() int {
	return len(a.Value)
}

func (a *JsonArray) Get(index int) *JsonElement {
	return &a.Value[index]
}

func (a *JsonArray) Add(element JsonElement) {
	a.Value = append(a.Value, element)
}

// type getters

func (a *JsonArray) GetArrayValue() *JsonArray {
	return a
}

func (a *JsonArray) GetBoolValue() *JsonBool {
	return nil
}

func (a *JsonArray) GetFloatValue() *JsonFloat {
	return nil
}

func (a *JsonArray) GetIntValue() *JsonInt {
	return nil
}

func (a *JsonArray) GetNullValue() *JsonNull {
	return nil
}

func (a *JsonArray) GetObjectValue() *JsonObject {
	return nil
}

func (a *JsonArray) GetStringValue() *JsonString {
	return nil
}
