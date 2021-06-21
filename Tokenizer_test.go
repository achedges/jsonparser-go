package jsonparser

import (
	"testing"
)

func TestJsonParser_TokenStream(t *testing.T) {
	json := `{ "string-key": "string-value", "int-key": 1, "float-key": 1.0, "bool-key": true, "null-key": null, "array-key": [ false ] }`

	parser := NewJsonParser(json)
	parser.Parse()

	expectedStream := []string{ "{", "ID: string-key", ":", "ID: string-value", ",", "ID: int-key", ":", "LIT: 1", ",", "ID: float-key", ":", "LIT: 1.0", ",", "ID: bool-key", ":", "LIT: true", ",", "ID: null-key", ":", "LIT: null", ",", "ID: array-key", ":", "[", "LIT: false", "]", "}" }

	if len(parser.TokenStream) != len(expectedStream) {
		t.Error("Unexpected token stream size")
		return
	}

	for index,token := range parser.TokenStream {
		tokenString := token.ToString()
		if tokenString != expectedStream[index] {
			t.Errorf("Unexpected token at index %d: %s should be %s", index, tokenString, expectedStream[index])
			return
		}
	}
}
