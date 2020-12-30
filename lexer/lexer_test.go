package lexer

import (
	"testing"
)

func TestLexer(t *testing.T) {

	input := `let five = 5;
let ten = 10;

let add = fn(x,y) {
  x + y;
};

let result = add(five, ten);
`

	cases := []struct {
		expectedType    string
		expectedLiteral string
	}{
		// let five = 5;
		{LET, "let"},
		{IDENT, "five"},
		{ASSIGN, "="},
		{INT, "5"},
		{SEMICOLON, ";"},

		// let ten = 10;
		{LET, "let"},
		{IDENT, "ten"},
		{ASSIGN, "="},
		{INT, "10"},
		{SEMICOLON, ";"},

		// let add = fn(x,y) {
		{LET, "let"},
		{IDENT, "add"},
		{ASSIGN, "="},
		{FUNCTION, "fn"},
		{LPAREN, "("},
		{IDENT, "x"},
		{COMMA, ","},
		{IDENT, "y"},
		{RPAREN, ")"},
		{LBRACE, "{"},

		//   x + y;
		{IDENT, "x"},
		{PLUS, "+"},
		{IDENT, "y"},
		{SEMICOLON, ";"},

		// };
		{RBRACE, "}"},
		{SEMICOLON, ";"},

		// let result = add(five, ten);
		{LET, "let"},
		{IDENT, "result"},
		{ASSIGN, "="},
		{IDENT, "add"},
		{LPAREN, "("},
		{IDENT, "five"},
		{COMMA, ","},
		{IDENT, "ten"},
		{RPAREN, ")"},
		{SEMICOLON, ";"},

		// EOF
		{EOF, ""},
		{EOF, ""}, // Repeat
	}

	l := New(input)

	for i, c := range cases {

		k := l.Token()

		if k.Type != c.expectedType {
			t.Fatalf("token %d: type wrong, expected %q, got %q", i, c.expectedType, k.Type)
		}

		if k.Literal != c.expectedLiteral {
			t.Fatalf("token %d: literal wrong", i)
		}

	}

}
