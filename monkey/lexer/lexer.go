package lexer

import "monkey/token"

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

/*
func New() Lexer → Returns a copy of the Lexer (less common for constructors).
func New() *Lexer → Returns a pointer (standard practice for struct constructors).
*/

// this is a package-level function that reates and returns a new *Lexer instance
func New(input string) *Lexer { // *Lexer means that the function returns a pointer to a Lexer struct (rather than a Lexer value itself)
	l := &Lexer{input: input} // creates a new Lexer struct instance and returns its memory address (a pointer to the struct)
	l.readChar()
	return l
}

/*
This is a method that operates on an existing Lexer instance
- the purpose of readChar is to give us the next character and advance our
position in the input string
  - first it checks whether we have reached the end of input; if that's the
    case, it sets l.ch to 0 (ASCII code for nul char)
  - if we havent reached the end yet, it sets l.ch to the next char by
    accessing l.input[l.readPosition]
  - finally, l.position is updated to l.readPosition and l.readPosition is
    incremented by one so that it always points to the next position we're
    going to read from and l.position always points to the position we
    last read
*/
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.NOT_EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	// adding this default branch to check for identifiers whenever l.ch
	// isn't one of our recognized chars
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

// helper function to peek ahead in the input and not move around in it
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

/*
  - in the nextToken method, we look at the current character under examination
    (l.ch) and return a token depending on which character it is
  - before returning the token we advance our pointers into the input so when
    we call NextToken() again the l.ch field is already updated
*/
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

/*
Notes on why readChar() and NextToken() are defined as methods on the Lexer
instance, but newToken is just a standalone function (not methods on the
token instance):

readChar and NextToken are Methods Because:
- They operate on Lexer state
- readChar() - Directly modifies the Lexer's internal state:

l.ch = ...          // Updates current character
l.position = ...    // Tracks position
l.readPosition++    // Moves read head
Needs access to Lexer's private fields (input, position, etc.)

- NextToken() - Consumes the Lexer's current character (l.ch)
- Advances the Lexer via l.readChar()
- Maintains parsing progress

They represent Lexer behavior
These are core actions the Lexer performs on itself to tokenize input.


newToken is a Standalone Function Because:
- It has no dependency on Lexer state
- Only needs the current character (ch byte) and token type
- Creates a new Token from scratch rather than modifying an existing one
- Pure function (same inputs always produce same outputs):
- No reliance on external state
- It's a utility function
- Like a factory that constructs Token objects
- Could be used independently (e.g., in tests or other packages)



Methods when the function:
- Needs access to receiver's state (Lexer's fields)
- Modifies the receiver
- Is part of the type's core behavior

Standalone functions when the function:
- Is stateless (only works with input args)
- Creates new instances rather than modifying existing ones
- Provides general utilities
*/
