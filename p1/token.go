package token

type TokenType string

const (
  ILLEGAL = "ILLEGAL"
  EOF = "EOF"
  IDENT = "IDENT"
  INT = "INT"

  ASSIGN = "="
  PLUS = "+"
  MINUS = "-"
  ASTERISK = "*"
  SLASH = "/"

  LT  = "<"
  GT =  ">"

  EQ = "=="
  NOT_EQ = "!="
  COMMA = ","
  SEMICOLON = ";"
  LPAREN = "("
  RPAREN = ")"
  LBRACE = "{"
  RBRACE = "}"
  FUNCTION = "FUNCTION"
  LET = "LET"
  TRUE = "TRUE"
  FALSE = "FALSE"
  IF = "IF"
  ELSE = "ELSE"
  RETURN = "RETURN"

  type Token string {
  
    Type TokenType
    Literal string
  }

  var keywords = map[string]TokenType {
    if tok, ok := Keywords[ident]; ok {
      return tok
    }
      return IDENT
  }


)