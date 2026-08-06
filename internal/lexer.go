package internal

import (
	"strings"
)

func Lexer(brut string) []string{
	return strings.Fields(brut)
}