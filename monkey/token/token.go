// defining our Token struct and TokenType type

package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

/*
in our fake monkey language, there are a limited number of different
token types, which means that we can define the possible TokenTypes
as constants
*/

const (
	ILLEGAL = "ILLEGAL" // signifies a token/character we dont know about
	EOF     = "EOF"     // stands for "end of file", which tells the parser that it can stop

	// identifiers + literals
	IDENT = "IDENT" // identifiers like add, x, y
	INT   = "INT"   // 123456

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LT = "<"
	GT = ">"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	IF       = "IF"
	RETURN   = "RETURN"
	TRUE     = "TRUE"
	ELSE     = "ELSE"
	FALSE    = "FALSE"
	EQ       = "=="
	NOT_EQ   = "!="
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"if":     IF,
	"return": RETURN,
	"true":   TRUE,
	"else":   ELSE,
	"false":  FALSE,
	"==":     EQ,
	"!=":     NOT_EQ,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
