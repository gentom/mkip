package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILEEGAL"
	EOF     = "EOF"

	// Identifier + Literal
	IDENT = "IDENT" // add , foobar, x, y, ...
	INT   = "INT"   // 1343456

	// Operator
	ASSIGN = "="
	PLUS   = "+"

	// Delimiter
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keyword
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
