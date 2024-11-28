package token

import "fmt"

type Type int

const (
	EOF Type = iota
	INTEGER
	PLUS
	MINUS
)

var TypeName = map[Type]string{
	EOF:     "EOF",
	INTEGER: "INTEGER",
	PLUS:    "PLUS",
	MINUS:   "MINUS",
}

var OperationValue = map[rune]Type{
	'+': PLUS,
	'-': MINUS,
}

type Token struct {
	Type  Type
	Value rune
}

func (t Token) String() string {
	return fmt.Sprintf("Token{Type: %s, Value: %s}", TypeName[t.Type], string(t.Value))
}

func NewToken(t Type, v rune) Token {
	return Token{Type: t, Value: v}
}
