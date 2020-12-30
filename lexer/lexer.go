package lexer

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

	defer l.readChar()

	switch l.Ch {
	case '=':
		return &Token{ASSIGN, string(l.Ch)}
	case ';':
		return &Token{SEMICOLON, string(l.Ch)}
	case '(':
		return &Token{LPAREN, string(l.Ch)}
	case ')':
		return &Token{RPAREN, string(l.Ch)}
	case ',':
		return &Token{COMMA, string(l.Ch)}
	case '+':
		return &Token{PLUS, string(l.Ch)}
	case '{':
		return &Token{LBRACE, string(l.Ch)}
	case '}':
		return &Token{RBRACE, string(l.Ch)}
	case 0:
		return &Token{EOF, ""}
	}

	return &Token{ILLEGAL, string(l.Ch)}
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
