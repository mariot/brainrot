package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

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

var ParsingInputError = fmt.Errorf("parsing input error")
var NotAnIntegerError = fmt.Errorf("not an integer error")
var IncorrectTypeError = fmt.Errorf("incorrect type error")

type Scanner struct {
	Text         string
	Position     int
	CurrentToken Token
}

func NewScanner(text string) Scanner {
	return Scanner{Text: text, Position: 0, CurrentToken: NewToken(EOF, '0')}
}

func (i *Scanner) error() {
	panic(ParsingInputError)
}

func (i *Scanner) getNextToken() Token {
	if i.Position > len(i.Text)-1 {
		return NewToken(EOF, '0')
	}

	currentChar := rune(i.Text[i.Position])

	if unicode.IsNumber(currentChar) {
		i.Position++
		return NewToken(INTEGER, currentChar)
	}

	if unicode.IsSpace(currentChar) {
		i.Position++
		return i.getNextToken()
	}

	if _, exists := OperationValue[currentChar]; exists {
		i.Position++
		return NewToken(OperationValue[currentChar], currentChar)
	}

	i.error()
	return NewToken(EOF, '0')
}

func (i *Scanner) eatIfType(tokenType Type) (Token, error) {
	if i.CurrentToken.Type == tokenType {
		return i.getNextToken(), nil
	}

	return NewToken(EOF, '0'), IncorrectTypeError
}

func (i *Scanner) eat(tokenType Type) {
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
		i.eat(INTEGER)
		if _, exists := OperationValue[i.CurrentToken.Value]; exists {
			break
		}
		leftValue += string(i.CurrentToken.Value)
	}

	op := i.CurrentToken
	i.eat(OperationValue[op.Value])

	rightValue := string(i.CurrentToken.Value)

	for {
		i.eat(INTEGER)
		if i.CurrentToken.Type == EOF {
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
	case PLUS:
		return leftInt + rightInt
	case MINUS:
		return leftInt - rightInt
	default:
		panic(IncorrectTypeError)
	}
}

func run(source string) {
	currentScanner := NewScanner(source)
	result := currentScanner.Expr()

	fmt.Println(result)
}

func runPrompt() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		if text == "exit" {
			break
		} else if text == "" {
			continue
		} else {
			run(text)
		}
	}
}

func main() {
	runPrompt()
}
