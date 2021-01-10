package parser

import (
	"monkey/ast"
	"monkey/lexer"
)

type Parser struct {
	lexer     *lexer.Lexer
	curToken  *lexer.Token // TODO: rename to "cur" ?
	peekToken *lexer.Token // TODO: rename to "peek" ?
}

func New(lexer *lexer.Lexer) *Parser {
	p := &Parser{
		lexer: lexer,
	}

	// TODO: fix ñap
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) ParseProgram() *ast.Program {

	program := &ast.Program{
		Statements: []interface{}{},
	}

	for p.curToken.Type != lexer.EOF {
		s := p.parseStatement()
		if s == nil {
			panic("nil statement!")
		}
		program.Statements = append(program.Statements, s)
		//p.nextToken()
	}

	return program
}

func (p *Parser) parseStatement() interface{} {

	switch p.curToken.Type {
	case lexer.LET:
		return p.parseLetStatement()
	case lexer.RETURN:
		return p.parseReturnStatement()
	default:
		return nil
	}
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.lexer.Token()
}

func (p *Parser) parseLetStatement() *ast.Let {

	if lexer.LET != p.curToken.Type {
		return nil // TODO: expected LET
	}
	p.nextToken() // consume 'let'

	s := &ast.Let{}

	if lexer.IDENT != p.curToken.Type {
		return nil // TODO: expected ident
	}

	if lexer.ASSIGN != p.peekToken.Type {
		return nil // TODO: expected assign
	}

	s.Name = &ast.Identifier{
		Token: p.curToken,
		Value: p.curToken.Literal, // Redundant?
	}

	// TODO: implement expression

	// Find semicolon
	for p.curToken.Type != lexer.SEMICOLON {
		p.nextToken()
		if p.curToken.Type == lexer.EOF {
			return nil // TODO: unexpected eof
		}
	}

	p.nextToken() // consume semicolon

	return s
}

func (p *Parser) parseReturnStatement() *ast.Return {

	if lexer.RETURN != p.curToken.Type {
		return nil // TODO: expected RETURN
	}

	s := &ast.Return{
		Token: p.curToken,
		Value: "",
	}

	p.nextToken() // consume 'return'

	// TODO: implement expression

	// Find semicolon
	for p.curToken.Type != lexer.SEMICOLON {
		p.nextToken()
		if p.curToken.Type == lexer.EOF {
			return nil // TODO: unexpected eof
		}
	}

	p.nextToken() // consume semicolon

	return s
}
