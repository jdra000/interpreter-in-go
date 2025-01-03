package token

type TokenType string

type Token struct {
	Type TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"fn": FUNCTION,
	"let": LET,
	"true": TRUE,
	"false": FALSE,
	"if": IF,
	"else": ELSE,
	"return": RETURN,
}

func LookupIdent(ident string) TokenType{
	if tok, ok := keywords[ident]; ok{
		return tok
	}
	return IDENT
}

const (
	// DIFFERENT ONES
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"

	// ASSIGNMENTS
	IDENT = "IDENT" // add, x, y, ...
	INT	= "INT" // 1234546..

	// OPERATORS
	ASSIGN = "="
	PLUS = "+"
	MINUS = "-"
	BANG = "!"
	ASTERISK = "*"
	SLASH = "/"
	LT = "<"
	GT = ">"

	// SYNTAX
	COMMA = ","
	SEMICOLON = ";"
	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"
	EQ = "=="
	NOT_EQ = "!="

	// KEYWORDS
	FUNCTION = "FUNCTION"
	LET = "LET"
	TRUE = "TRUE"
	FALSE = "FALSE"
	IF = "IF"
	ELSE = "ELSE"
	RETURN = "RETURN"
)
