package token

type TokenType string

type Token struct {
	Type TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"fn": FUNCTION,
	"let": LET,
}

func LookupIdent(ident string) TokenType{
	if tok, ok := keywords[ident]; ok{
		return tok
	}
	return IDENT
}

const (
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"

	IDENT = "IDENT" // add, x, y, ...
	INT	= "INT" // 1234546..

	ASSIGN = "="
	PLUS = "+"

	COMMA = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	FUNCTION = "FUNCTION"
	LET = "LET"
)
