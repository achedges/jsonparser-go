package types

import "fmt"

type JsonInt struct {
	 Value int64
}

func (i *JsonInt) Serialize(indent int) string {
	_ = indent
	return fmt.Sprint(i.Value)
}

func (i *JsonInt) GetArrayValue() *JsonArray {
	return nil
}

func (i *JsonInt) GetBoolValue() *JsonBool {
	return nil
}

func (i *JsonInt) GetFloatValue() *JsonFloat {
	return &JsonFloat{float64(i.Value) }
}

func (i *JsonInt) GetIntValue() *JsonInt {
	return i
}

func (i *JsonInt) GetNullValue() *JsonNull {
	return nil
}

func (i *JsonInt) GetObjectValue() *JsonObject {
	return nil
}

func (i *JsonInt) GetStringValue() *JsonString {
	return nil
}
