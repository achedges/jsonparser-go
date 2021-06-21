package types

type JsonNull struct {

}

func (n *JsonNull) Serialize(indent int) string {
	_ = indent
	return "null"
}

// type getters

func (n *JsonNull) GetArrayValue() *JsonArray {
	return nil
}

func (n *JsonNull) GetBoolValue() *JsonBool {
	return nil
}

func (n *JsonNull) GetFloatValue() *JsonFloat {
	return nil
}

func (n *JsonNull) GetIntValue() *JsonInt {
	return nil
}

func (n *JsonNull) GetNullValue() *JsonNull {
	return n
}

func (n *JsonNull) GetObjectValue() *JsonObject {
	return nil
}

func (n *JsonNull) GetStringValue() *JsonString {
	return nil
}
