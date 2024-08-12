package token

import "testing"

func TestTokenStruct(t *testing.T) {
	_ = Token{Type: "IDENT", Literal: "foobar"}
	_ = Token{Type: "INT", Literal: "123456"}
	_ = Token{Type: "ASSIGN", Literal: "="}
	_ = Token{Type: "PLUS", Literal: "+"}
	_ = Token{Type: "COMMA", Literal: ","}
	_ = Token{Type: "SEMICOLON", Literal: ";"}
	_ = Token{Type: "LPAREN", Literal: "("}
	_ = Token{Type: "RPAREN", Literal: ")"}
	_ = Token{Type: "LBRACE", Literal: "{"}
	_ = Token{Type: "RBRACE", Literal: "}"}
}
