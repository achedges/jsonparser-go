package jsonparser

import "testing"

func TestArrayParser(t *testing.T) {
	jsonArray := `[ "string value", 33, 5.2, true, false, null, { "object key": "object value" } ]`
	parser := NewJsonParser(jsonArray)
	array := parser.Parse().GetArrayValue()

	if array == nil {
		t.Error("Result is not a JSON array")
		return
	}

	if array.GetSize() != 7 {
		t.Errorf("Unexpected array size: %d", array.GetSize())
		return
	}

	if (*array.Get(0)).GetStringValue().Value != "string value" {
		t.Error("Unexpected string value")
		return
	}

	if (*array.Get(1)).GetIntValue().Value != 33 {
		t.Error("Unexpected int value")
		return
	}

	if (*array.Get(2)).GetFloatValue().Value != 5.2 {
		t.Error("Unexpected float value")
		return
	}

	if !(*array.Get(3)).GetBoolValue().Value {
		t.Error("Unexpected bool value")
		return
	}

	if (*array.Get(4)).GetBoolValue().Value {
		t.Error("Unexpected bool value")
		return
	}

	if (*array.Get(5)).GetNullValue() == nil {
		t.Error("NULL value not found")
		return
	}

	if (*array.Get(6)).GetObjectValue() == nil {
		t.Error("Object value not found")
		return
	}

	obj := (*array.Get(6)).GetObjectValue()
	if (*obj.Get("object key")).GetStringValue().Value != "object value" {
		t.Error("Unexpected object key-value pair")
		return
	}
}

func TestObjectParser(t *testing.T) {
	input := `{ "string": "value", "int": 12, "float": 77.2, "true": true, "false": false, "null": null, "array": [ 1, 2 ], "object": { "nested": "value" } }`
	parser := NewJsonParser(input)
	object := parser.Parse().GetObjectValue()

	if object == nil {
		t.Error("Resulting is not a JSON object")
		return
	}

	jsonString := (*object.Get("string")).GetStringValue()
	jsonInt := (*object.Get("int")).GetIntValue()
	jsonFloat := (*object.Get("float")).GetFloatValue()
	jsonTrue := (*object.Get("true")).GetBoolValue()
	jsonFalse := (*object.Get("false")).GetBoolValue()
	jsonNull := (*object.Get("null")).GetNullValue()
	jsonArray := (*object.Get("array")).GetArrayValue()
	jsonObject := (*object.Get("object")).GetObjectValue()

	if jsonString == nil || jsonString.Value != "value" {
		t.Error("Unexpected string value")
		return
	}

	if jsonInt == nil || jsonInt.Value != 12 {
		t.Error("Unexpected int value")
		return
	}

	if jsonFloat == nil || jsonFloat.Value != 77.2 {
		t.Error("Unexpected float value")
		return
	}

	if jsonTrue == nil || !jsonTrue.Value {
		t.Error("Unexpected bool value")
		return
	}

	if jsonFalse == nil || jsonFalse.Value {
		t.Error("Unexpected bool value")
		return
	}

	if jsonNull == nil {
		t.Error("JsonNull element not found")
		return
	}

	if jsonArray == nil || jsonArray.GetSize() != 2 {
		t.Error("Array not found, or unexpected size")
		return
	}

	if jsonObject == nil || jsonObject.GetSize() != 1 {
		t.Error("Object not found, or unexpected size")
		return
	}

	jsonNestedObject := (*jsonObject.Get("nested")).GetStringValue()
	if jsonNestedObject == nil || jsonNestedObject.Value != "value" {
		t.Error("Nested object not found, or unexpected value")
		return
	}
}
