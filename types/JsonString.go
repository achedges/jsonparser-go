package types

type JsonString struct {
	Value string
}

func (s *JsonString) Serialize(indent int) string {
	_ = indent
	return "\"" + s.Value + "\""
}

// type getters

func (s *JsonString) GetArrayValue() *JsonArray {
	return nil
}

func (s *JsonString) GetBoolValue() *JsonBool {
	return nil
}

func (s *JsonString) GetFloatValue() *JsonFloat {
	return nil
}

func (s *JsonString) GetIntValue() *JsonInt {
	return nil
}

func (s *JsonString) GetNullValue() *JsonNull {
	return nil
}

func (s *JsonString) GetObjectValue() *JsonObject {
	return nil
}

func (s *JsonString) GetStringValue() *JsonString {
	return s
}
