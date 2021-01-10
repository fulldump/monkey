package ast

import (
	"monkey/lexer"
)

type Let struct {
	Token *lexer.Token
	Name  *Identifier
	Value *Expression
}
