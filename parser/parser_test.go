package parser

import (
	"monkey/ast"
	"monkey/lexer"
	"testing"
)

type expected struct {
	identifier string
}

func TestLetStatements(t *testing.T) {

	input := `
let x = 1;
let y = 2;
let z = 3;
`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	if program == nil {
		t.Fatalf("parseprogram return nil")
	}
	if len(program.Statements) != 3 {
		t.Fatal("expected 3 statemtns")
	}

	cases := []*expected{
		{identifier: "x"},
		{identifier: "y"},
		{identifier: "z"},
	}

	for i, expected := range cases {
		s := program.Statements[i]

		assertStatement(t, s, expected)

	}
}

func assertStatement(t *testing.T, statement interface{}, e *expected) {

	switch s := statement.(type) {
	case *ast.Let:
		if s.Name.Value != e.identifier {
			t.Fatalf("unexpected identifier '%s', expected '%s'", s.Name.Value, e.identifier)
		}
	default:
		t.Fatalf("unexpected statement")
	}

}
