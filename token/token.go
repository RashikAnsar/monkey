package token

// TokenType allows us to distinguish between different types of tokens.
type TokenType string

const (
	// ILLEGAL signifies a token/character we don’t know about
	ILLEGAL = "ILLEGAL"

	// EOF stands for “end of file”, which tells our parser later on that it can stop
	EOF = "EOF"

	// IDENT Identifiers
	IDENT = "IDENT" // add, foobar, x, y, ...
	// INT literals
	INT = "INT" // 1343456

	// ********************** OPERATORS ***********************

	// ASSIGN Operator
	ASSIGN = "="
	// PLUS Operator
	PLUS = "+"
	// MINUS Operator
	MINUS = "-"
	// BANG Operator
	BANG = "!"
	// ASTERISK Operator
	ASTERISK = "*"
	// SLASH Operator
	SLASH = "/"
	// LT Operator
	LT = "<"
	// GT Operator
	GT = ">"

	// ********************* DELIMITERS ************************

	// COMMA Delimiters
	COMMA = ","
	// SEMICOLON Delimiters
	SEMICOLON = ";"
	// LPAREN Delimiters
	LPAREN = "("
	// RPAREN Delimiters
	RPAREN = ")"
	// LBRACE Delimiters
	LBRACE = "{"
	// RBRACE Delimiters
	RBRACE = "}"

	// ******************** KEYWORDS ****************************

	// FUNCTION Keyword
	FUNCTION = "FUNCTION"
	// LET Keyword
	LET = "LET"
)

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

// Token contains Type and Description
type Token struct {
	Type    TokenType
	Literal string
}

// LookupIdent look up the identifier from keywords
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
