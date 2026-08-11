package internal

import "github.com/google/shlex"

func Lexer(brut string) []string {
	tokens, err := shlex.Split(brut)
	if err != nil {
		return nil
	}

	return tokens
}