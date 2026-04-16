package helpers

import (
	"strconv"
)

func ApplyHex(tokens []string) []string {
	for i := 0; i < len(tokens); i++ {
		if tokens[i] == "(hex)" && i > 0{
			n, _ := strconv.ParseInt(tokens[i-1], 16, 64)
			tokens[i-1] = strconv.FormatInt(n, 10)
			tokens[i] = ""
		}
	}
	return tokens
}

func ApplyBin(tokens []string) []string {
	for i := 0; i < len(tokens); i++ {
		if tokens[i] == "(bin)" {
			n, _ := strconv.ParseInt(tokens[i-1], 2, 64)
			tokens[i-1] = strconv.FormatInt(n, 10)
			tokens[i] = ""
		}
	}
	return tokens
}