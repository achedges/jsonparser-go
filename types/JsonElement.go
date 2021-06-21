package types

type JsonElement interface {
	Serialize(indent int) string
	GetArrayValue() *JsonArray
	GetBoolValue() *JsonBool
	GetFloatValue() *JsonFloat
	GetIntValue() *JsonInt
	GetNullValue() *JsonNull
	GetObjectValue() *JsonObject
	GetStringValue() *JsonString
}
