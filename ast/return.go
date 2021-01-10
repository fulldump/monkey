package ast

import (
	"monkey/lexer"
)

type Return struct {
	Token *lexer.Token
	Value Expression
}
