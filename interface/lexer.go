package _interface

type Lexer interface {
	New(input string) *Lexer
	readChar()
	NextToken()
}
