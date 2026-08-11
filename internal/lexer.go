package internal

import (
	"strings"
)

func Lexer(brut string) []string {
	res := []string{}
	var buf strings.Builder
	hasContent := false

	i := 0
	n := len(brut)
	for i < n {
		c := brut[i]

		switch c {
		case ' ':
			if hasContent {
				res = append(res, buf.String())
				buf.Reset()
				hasContent = false
			}
			i++

		case '\\':
			if i+1 < n {
				buf.WriteByte(brut[i+1])
				hasContent = true
				i += 2
			} else {
				buf.WriteByte(c)
				hasContent = true
				i++
			}

		case '\'':
			relEnd := strings.Index(brut[i+1:], "'")
			if relEnd == -1 {
				buf.WriteString(brut[i+1:])
				hasContent = true
				i = n
				continue
			}
			end := i + 1 + relEnd
			buf.WriteString(brut[i+1 : end])
			hasContent = true
			i = end + 1

		case '"':
			j := i + 1
			for j < n {
				if brut[j] == '"' {
					break
				}
				if brut[j] == '\\' && j+1 < n && isEscapableInDouble(brut[j+1]) {
					buf.WriteByte(brut[j+1])
					j += 2
					continue
				}
				buf.WriteByte(brut[j])
				j++
			}
			hasContent = true
			if j < n {
				i = j + 1 
			} else {
				i = j 
			}

		default:
			buf.WriteByte(c)
			hasContent = true
			i++
		}
	}

	if hasContent {
		res = append(res, buf.String())
	}
	return res
}

func isEscapableInDouble(c byte) bool {
	switch c {
	case '$', '`', '"', '\\', '\n':
		return true
	}
	return false
}
