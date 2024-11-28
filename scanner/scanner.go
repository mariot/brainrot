package scanner

import (
	"fmt"
	"github.com/mariot/token"
	"strconv"
	"unicode"
)

var ParsingInputError = fmt.Errorf("parsing input error")
var NotAnIntegerError = fmt.Errorf("not an integer error")
var IncorrectTypeError = fmt.Errorf("incorrect type error")

type Scanner struct {
	Text         string
	Position     int
	CurrentToken token.Token
}

func NewScanner(text string) Scanner {
	return Scanner{Text: text, Position: 0, CurrentToken: token.NewToken(token.EOF, '0')}
}

func (i *Scanner) error() {
	panic(ParsingInputError)
}

func (i *Scanner) getNextToken() token.Token {
	if i.Position > len(i.Text)-1 {
		return token.NewToken(token.EOF, '0')
	}

	currentChar := rune(i.Text[i.Position])

	if unicode.IsNumber(currentChar) {
		i.Position++
		return token.NewToken(token.INTEGER, currentChar)
	}

	if unicode.IsSpace(currentChar) {
		i.Position++
		return i.getNextToken()
	}

	if _, exists := token.OperationValue[currentChar]; exists {
		i.Position++
		return token.NewToken(token.OperationValue[currentChar], currentChar)
	}

	i.error()
	return token.NewToken(token.EOF, '0')
}

func (i *Scanner) eatIfType(tokenType token.Type) (token.Token, error) {
	if i.CurrentToken.Type == tokenType {
		return i.getNextToken(), nil
	}

	return token.NewToken(token.EOF, '0'), IncorrectTypeError
}

func (i *Scanner) eat(tokenType token.Type) {
	if i.CurrentToken.Type == tokenType {
		i.CurrentToken = i.getNextToken()
	} else {
		i.error()
	}
}

func (i *Scanner) Expr() int {
	i.CurrentToken = i.getNextToken()

	leftValue := string(i.CurrentToken.Value)

	for {
		i.eat(token.INTEGER)
		if _, exists := token.OperationValue[i.CurrentToken.Value]; exists {
			break
		}
		leftValue += string(i.CurrentToken.Value)
	}

	op := i.CurrentToken
	i.eat(token.OperationValue[op.Value])

	rightValue := string(i.CurrentToken.Value)

	for {
		i.eat(token.INTEGER)
		if i.CurrentToken.Type == token.EOF {
			break
		}
		rightValue += string(i.CurrentToken.Value)
	}

	leftInt, err := strconv.Atoi(leftValue)
	if err != nil {
		panic(NotAnIntegerError)
	}

	rightInt, err := strconv.Atoi(rightValue)
	if err != nil {
		panic(NotAnIntegerError)
	}

	switch op.Type {
	case token.PLUS:
		return leftInt + rightInt
	case token.MINUS:
		return leftInt - rightInt
	default:
		panic(IncorrectTypeError)
	}
}
