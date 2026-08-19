package internal

type TokenType string


//value of tokentype that we are using
const (
	Word        TokenType = "WORD"
	Pipe        TokenType = "PIPE"
	RedirectOut TokenType = "REDIRECT_OUT"
	RedirectAppend TokenType = "RADIRECT_APPEND"
)

//struct token 
type Token struct {
	Value string
	Type  string
}
