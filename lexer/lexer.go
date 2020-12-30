package lexer

import (
	"strings"
	"unicode"
)

type Lexer struct {
	Input        string // TODO: change to reader
	Position     int
	ReadPosition int
	Ch           byte // TODO: change to rune
}

func New(input string) *Lexer {
	l := &Lexer{
		Input: input,
	}
	l.readChar()
	return l
}

func (l *Lexer) Token() *Token {

	l.skipWhitespaces()

	t := &Token{
		Literal: string(l.Ch), // Fill default token
	}

	switch l.Ch {
	case '=':
		if l.peekChar() == '=' {
			t.Type = EQ
			t.Literal = "=="
			l.readChar()
			l.readChar()
			return t
		}
		t.Type = ASSIGN
		l.readChar()
		return t
	case ';':
		t.Type = SEMICOLON
		l.readChar()
		return t
	case '(':
		t.Type = LPAREN
		l.readChar()
		return t
	case ')':
		t.Type = RPAREN
		l.readChar()
		return t
	case ',':
		t.Type = COMMA
		l.readChar()
		return t
	case '+':
		t.Type = PLUS
		l.readChar()
		return t
	case '-':
		t.Type = MINUS
		l.readChar()
		return t
	case '/':
		t.Type = SLASH
		l.readChar()
		return t
	case '*':
		t.Type = ASTERISK
		l.readChar()
		return t
	case '!':
		if l.peekChar() == '=' {
			t.Type = NEQ
			t.Literal = "!="
			l.readChar()
			l.readChar()
			return t
		}
		t.Type = BANG
		l.readChar()
		return t
	case '<':
		t.Type = LT
		l.readChar()
		return t
	case '>':
		t.Type = GT
		l.readChar()
		return t
	case '{':
		t.Type = LBRACE
		l.readChar()
		return t
	case '}':
		t.Type = RBRACE
		l.readChar()
		return t
	case 0:
		return &Token{EOF, ""}
	}

	if unicode.IsLetter(int32(l.Ch)) {
		literal := l.readIdentifier()
		tokenType := IDENT

		k := strings.ToLower(literal)
		if tt, isKeyword := keywords[k]; isKeyword {
			tokenType = tt
		}

		return &Token{tokenType, literal}
	}

	if isDigit(l.Ch) {
		literal := l.readNumber()
		return &Token{INT, literal}
	}

	return &Token{ILLEGAL, string(l.Ch)}
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) readIdentifier() string {
	position := l.Position
	for unicode.IsLetter(int32(l.Ch)) {
		l.readChar()
	}
	return l.Input[position:l.Position]
}

func (l *Lexer) readChar() {
	if l.ReadPosition >= len(l.Input) {
		l.Ch = 0
	} else {
		l.Ch = l.Input[l.ReadPosition]
	}
	l.Position = l.ReadPosition
	l.ReadPosition++
}

func (l *Lexer) skipWhitespaces() {
	for l.Ch == ' ' || l.Ch == '\t' || l.Ch == '\n' || l.Ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readNumber() string {
	position := l.Position
	for isDigit(l.Ch) {
		l.readChar()
	}
	return l.Input[position:l.Position]
}

func (l *Lexer) peekChar() byte {
	if l.ReadPosition >= len(l.Input) {
		return 0
	}
	return l.Input[l.ReadPosition]
}
