package lexer

import "github.com/ssttuu/monkey/pkg/token"

func New(input string) *Lexer {
	l := &Lexer{input:input}
	l.readChar()
	return l
}

type Lexer struct {
	input string
	position int
	readPosition int
	ch byte
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.Equal, Literal: literal}
		} else {
			tok = newToken(token.Assign, l.ch)
		}
	case '+':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.PlusEqual, Literal: literal}
		} else {
			tok = newToken(token.Plus, l.ch)
		}
	case '-':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.MinusEqual, Literal: literal}
		} else {
			tok = newToken(token.Minus, l.ch)
		}
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.NotEqual, Literal: literal}
		} else {
			tok = newToken(token.Not, l.ch)
		}
	case '/':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.DivideEqual, Literal: literal}
		} else {
			tok = newToken(token.Divide, l.ch)
		}
	case '*':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.MultiplyEqual, Literal: literal}
		} else {
			tok = newToken(token.Multiply, l.ch)
		}

	case '<':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.LessThanOrEqual, Literal: literal}
		} else {
			tok = newToken(token.LessThan, l.ch)
		}
	case '>':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.GreaterThanOrEqual, Literal: literal}
		} else {
			tok = newToken(token.GreaterThan, l.ch)
		}

	case ';':
		tok = newToken(token.Semicolon, l.ch)
	case ',':
		tok = newToken(token.Comma, l.ch)

	case '(':
		tok = newToken(token.LeftParentheses, l.ch)
	case ')':
		tok = newToken(token.RightParentheses, l.ch)

	case '{':
		tok = newToken(token.LeftBrace, l.ch)
	case '}':
		tok = newToken(token.RightBrace, l.ch)

	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdentifier(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.Integer
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.Illegal, l.ch)
		}
	}
	l.readChar()
	return tok
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func newToken(t token.TokenType, ch byte) token.Token {
	return token.Token{Type: t, Literal:string(ch)}
}