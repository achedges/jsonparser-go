package types

import (
	"testing"
)

func TestJsonArray_Serialize(t *testing.T) {
	expected := "[ true, false ]"
	a := NewJsonArray()
	a.Add(&JsonBool{ Value: true })
	a.Add(&JsonBool{ Value: false })

	actual := a.Serialize(0)

	if actual != expected {
		t.Fail()
	}
}

func TestJsonBool_Serialize(t *testing.T) {
	trueVal := JsonBool{ Value: true }
	if trueVal.Serialize(0) != "true" {
		t.Fail()
	}

	falseVal := JsonBool{ Value: false }
	if falseVal.Serialize(0) != "false" {
		t.Fail()
	}
}

func TestJsonFloat_Serialize(t *testing.T) {
	floatVal := JsonFloat{ Value: 12.34 }
	if floatVal.Serialize(0) != "12.34" {
		t.Fail()
	}
}

func TestJsonInt_Serialize(t *testing.T) {
	intVal := JsonInt{ Value: 123 }
	if intVal.Serialize(0) != "123" {
		t.Fail()
	}
}

func TestJsonNull_Serialize(t *testing.T) {
	nullVal := JsonNull{}
	if nullVal.Serialize(0) != "null" {
		t.Fail()
	}
}

func TestJsonObject_Serialize(t *testing.T) {
	objVal := NewJsonObject()
	objVal.Set("some-string", &JsonString{ Value: "some-value" })
	objVal.Set("some-int", &JsonInt{ Value: 17 })
	objVal.Set("some-float", &JsonFloat{ Value: 8.56 })
	objVal.Set("some-bool", &JsonBool{ Value: true })
	objVal.Set("some-null", &JsonNull{})

	arrVal := NewJsonArray()
	arrVal.Add(&JsonInt{ Value: 12 })
	arrVal.Add(&JsonInt{ Value: 13 })

	innerObj := NewJsonObject()
	innerObj.Set("inner-string", &JsonString{ Value: "inner-value" })

	objVal.Set("some-array", arrVal)
	objVal.Set("some-object", innerObj)

	expected := "{ \"some-array\": [ 12, 13 ], \"some-bool\": true, \"some-float\": 8.56, \"some-int\": 17, \"some-null\": null, \"some-object\": { \"inner-string\": \"inner-value\" }, \"some-string\": \"some-value\" }"
	actual := objVal.Serialize(0)

	if actual != expected {
		t.Fail()
	}
}

func TestJsonString_Serialize(t *testing.T) {
	stringVal := JsonString{ Value: "testing string value" }
	if stringVal.Serialize(0) != "\"testing string value\"" {
		t.Fail()
	}
}
