package token

type TokenType string

type Token struct {

  Type TokenType
  literal string

}

const (
  ILLEGAL = "ILLEGAL"
  EOF = "EOF"
  IDENT = "IDENT"
  INT = "INT"
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