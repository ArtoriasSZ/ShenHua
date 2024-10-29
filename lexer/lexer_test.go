package lexer

import (
	"fmt"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `let five = 5;let ten = 10;let add = fn(x, y) {x + y;};let result = add(five, ten);`
	l := New(input)
	for {
		tok := l.NextToken()
		fmt.Println(tok)
	}
}
