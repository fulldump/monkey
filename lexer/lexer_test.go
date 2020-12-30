package lexer

import (
	"testing"
)

func TestLexer(t *testing.T) {

	input := `=+(){},;`

	cases := []struct {
		expectedType    string
		expectedLiteral string
	}{
		{ASSIGN, "="},
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
