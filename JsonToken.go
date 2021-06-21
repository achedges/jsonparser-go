package jsonparser

type JsonToken struct {
	Token string
	Text string
}

func NewJsonToken(token string, text string) JsonToken {
	return JsonToken { Token: token, Text: text }
}

func (t *JsonToken) ToString() string {
	if t.Token == t.Text {
		return t.Token
	} else {
		return t.Token + ": " + t.Text
	}
}
