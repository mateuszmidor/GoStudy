package yaml

import (
	"fmt"
	"text/scanner"
)

type lexer struct {
	scan scanner.Scanner
}

func (lex *lexer) nextToken() string {
	token := lex.scan.Scan()
	text := lex.scan.TokenText()
	fmt.Printf("[lexer] Fetched token: %s[%s]\n", tokenToString(token), text)
	return text
}

func (lex *lexer) nextRune() rune {
	token := lex.scan.Next()
	fmt.Printf("[lexer] Fetched rune: %q\n", tokenToString(token))
	return token
}

func (lex *lexer) consumeSpaces(count int) bool {
	for i := 0; i < count; i++ {
		if lex.nextRune() != ' ' {
			return false
		}
	}

	if lex.scan.Peek() == ' ' {
		panic(fmt.Sprintf("More spaces than expected"))
	}

	// managed to read all the requested spaces
	return true
}

func (lex *lexer) consumeRunes(want string) {
	for _, wantRune := range want {
		actualRune := lex.nextRune()
		if actualRune != wantRune {
			panic(fmt.Sprintf("got %q, expected: %q", actualRune, wantRune))
		}
	}
}

func (lex *lexer) consumeEOL() {
	r := lex.nextRune()
	if r != '\n' && r != scanner.EOF {
		panic(fmt.Sprintf("got %q, expected: end of line or end of file", r))
	}
}

// "Person:\n"
func (lex *lexer) readObject() string {
	key := lex.nextToken()
	lex.consumeRunes(":")
	lex.consumeEOL()
	return key
}

// "Flag: true\n"
func (lex *lexer) readKeyValue() (string, string) {
	key := lex.nextToken()
	lex.consumeRunes(": ")
	value := lex.nextToken()
	lex.consumeEOL()
	return key, value
}

// "- 24\n"
func (lex *lexer) readListItem() (string, bool) {
	nextrune := lex.scan.Peek()
	if nextrune != '-' {
		return "", false
	}

	lex.consumeRunes("- ")
	value := lex.nextToken()
	lex.consumeEOL()
	return value, true
}

func tokenToString(r rune) string {
	switch r {
	case scanner.String:
		return "string"
	case scanner.Ident:
		return "identifier"
	case scanner.Int:
		return "integer"
	case scanner.EOF:
		return "End Of File"
	default:
		return fmt.Sprintf("%c", r)
	}
}
