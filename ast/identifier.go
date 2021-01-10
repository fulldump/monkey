package ast

import (
	"monkey/lexer"
)

type Identifier struct {
	Token *lexer.Token
	Value string
}
