package token

type TokenType string

const (
	ILLEGAL  TokenType = "ILLEGAL"
	EOF      TokenType = "EOF"
	IDENT    TokenType = "IDENT"
	INT      TokenType = "INT"
	STRING   TokenType = "STRING"
	TRUE     TokenType = "TRUE"
	FALSE    TokenType = "FALSE"
	NULL     TokenType = "NULL"
	LBRACE   TokenType = "{"
	RBRACE   TokenType = "}"
	LBRACKET TokenType = "["
	RBRACKET TokenType = "]"
	COLON    TokenType = ":"
	COMMA    TokenType = ","
)

type Token struct {
	Type    TokenType
	Literal string
}

func LookupIdent(ident string) TokenType {
	switch ident {
	case "true":
		return TRUE
	case "false":
		return FALSE
	case "null":
		return NULL
	default:
		return IDENT
	}
}
