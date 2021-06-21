package types

import "fmt"

type JsonFloat struct {
	Value float64
}

func (f *JsonFloat) Serialize(indent int) string {
	_ = indent
	return fmt.Sprint(f.Value)
}

func (f *JsonFloat) GetArrayValue() *JsonArray {
	return nil
}

func (f *JsonFloat) GetBoolValue() *JsonBool {
	return nil
}

func (f *JsonFloat) GetFloatValue() *JsonFloat {
	return f
}

func (f *JsonFloat) GetIntValue() *JsonInt {
	return &JsonInt{int64(f.Value) }
}

func (f *JsonFloat) GetNullValue() *JsonNull {
	return nil
}

func (f *JsonFloat) GetObjectValue() *JsonObject {
	return nil
}

func (f *JsonFloat) GetStringValue() *JsonString {
	return nil
}
