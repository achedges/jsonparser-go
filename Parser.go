package jsonparser

import (
	"fmt"
	"jsonparser/types"
	"strconv"
	"strings"
)

type JsonParser struct {
	input string
	n     int
	i     int

	TokenStream []JsonToken
}

func NewJsonParser(inputStr string) JsonParser {
	return JsonParser {
		input:       inputStr,
		n:           len(inputStr),
		i:           0,
		TokenStream: make([]JsonToken, 0, 20),
	}
}

func (p *JsonParser) Parse() types.JsonElement {
	p.tokenize()
	p.i = 0

	if p.TokenStream[0].Token == "{" {
		return p.parseObject()
	} else if p.TokenStream[0].Token == "[" {
		return p.parseArray()
	} else {
		fmt.Println("JSON documents must start with '{' or '['")
		return nil
	}
}

func (p *JsonParser) tokenize() {
	p.i = 0
	for p.i < p.n {
		if isEmptyChar(p.input[p.i]) { // discard?
			p.i += 1
			continue
		} else if isJsonToken(p.input[p.i]) { // punctuation/token literal?
			p.TokenStream = append(p.TokenStream, NewJsonToken(string(p.input[p.i]), string(p.input[p.i])))
		} else if p.input[p.i] == '"' { // identifier?
			var id string
			p.i += 1
			for p.i < p.n {
				if p.input[p.i] == '"' && p.input[p.i- 1] != '\\' {
					break
				} else if p.input[p.i] == '"' && p.input[p.i- 1] == '\\' {
					id = strings.TrimSuffix(id, string('\\'))
					id += string('"')
				} else {
					id += string(p.input[p.i])
				}

				p.i += 1
			}

			p.TokenStream = append(p.TokenStream, NewJsonToken("ID", id))
		} else { // must be a literal
			var lit string
			for p.i < p.n {
				if p.input[p.i] == ',' || p.input[p.i] == ']' || p.input[p.i] == '}' {
					p.i -= 1 // back up so we can capture the ending token in the stream
					break
				} else if !isEmptyChar(p.input[p.i]) {
					lit += string(p.input[p.i])
				}

				p.i += 1
			}

			p.TokenStream = append(p.TokenStream, NewJsonToken("LIT", lit))
		}

		p.i += 1
	}
}

func (p *JsonParser) nextToken() JsonToken {
	p.i += 1
	if p.i < len(p.TokenStream) {
		return p.TokenStream[p.i]
	} else {
		return NewJsonToken("EOF", "EOF")
	}
}

func (p *JsonParser) parseLiteral(text string) types.JsonElement {
	if text == "true" {
		return &types.JsonBool { Value: true }
	} else if text == "false" {
		return &types.JsonBool { Value: false }
	} else if text == "null" {
		return &types.JsonNull {}
	} else if strings.Index(text, ".") != -1 {
		f,_ := strconv.ParseFloat(text, 64)
		return &types.JsonFloat { Value: f }
	} else {
		i,_ := strconv.ParseInt(text, 10, 64)
		return &types.JsonInt { Value: i }
	}
}

func (p *JsonParser) parseArray() types.JsonElement {
	array := types.NewJsonArray()
	token := p.nextToken()

	for token.Token != "]" {
		if token.Token == "ID" {
			array.Add(&types.JsonString { Value: token.Text })
		} else if token.Token == "LIT" {
			array.Add(p.parseLiteral(token.Text))
		} else if token.Token == "[" {
			array.Add(p.parseArray())
		} else if token.Token == "{" {
			array.Add(p.parseObject())
		}

		token = p.nextToken()
		if token.Token == "EOF" {
			break
		}
	}

	return array
}

func (p *JsonParser) parseObject() types.JsonElement {
	object := types.NewJsonObject()
	key := p.nextToken()

	for key.Token != "}" {
		if key.Token == "," {
			key = p.nextToken()
		}

		separator := p.nextToken()
		value := p.nextToken()

		if key.Token != "ID" {
			fmt.Println("Invalid token detected for object key: " + key.Token)
			return object
		}

		if separator.Token != ":" {
			fmt.Println("Invalid token detected for key-value separator: " + separator.Token)
			return object
		}

		if value.Token == "{" {
			object.Set(key.Text, p.parseObject())
		} else if value.Token == "[" {
			object.Set(key.Text, p.parseArray())
		} else if value.Token == "ID" {
			object.Set(key.Text, &types.JsonString { Value: value.Text })
		} else if value.Token == "LIT" {
			object.Set(key.Text, p.parseLiteral(value.Text))
		} else {
			fmt.Println("Unknown value token: " + value.Token)
			return object
		}

		key = p.nextToken()
	}

	return object
}

func isEmptyChar(c byte) bool {
	switch c {
	case ' ':
		return true
	case '\t':
		return true
	case '\n':
		return true
	case '\r':
		return true
	}

	return false
}

func isJsonToken(c byte) bool {
	switch c {
	case '{':
		return true
	case '}':
		return true
	case '[':
		return true
	case ']':
		return true
	case ':':
		return true
	case ',':
		return true
	}

	return false
}
