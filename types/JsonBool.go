package types

type JsonBool struct {
	Value bool
}

func (b *JsonBool) Serialize(indent int) string {
	_ = indent // not needed
	if b.Value {
		return "true"
	} else {
		return "false"
	}
}

func (b *JsonBool) GetArrayValue() *JsonArray {
	return nil
}

func (b *JsonBool) GetBoolValue() *JsonBool {
	return b
}

func (b *JsonBool) GetFloatValue() *JsonFloat {
	return nil
}

func (b *JsonBool) GetIntValue() *JsonInt {
	return nil
}

func (b *JsonBool) GetNullValue() *JsonNull {
	return nil
}

func (b *JsonBool) GetObjectValue() *JsonObject {
	return nil
}

func (b *JsonBool) GetStringValue() *JsonString {
	return nil
}