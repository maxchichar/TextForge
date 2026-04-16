package helpers

import(
	"strings"
)

func ApplyAn(tokens []string) []string {
	vowel := "aeiouhAEIOUH"
	for i := 0; i < len(tokens); i++ {
		if tokens[i] == "a" || tokens[i] == "A" && strings.ContainsAny(tokens[i+1], vowel) {
			switch tokens[i] {
			case "a":
				tokens[i] = "an"
			case "A":
				tokens[i] = "An"
			}
		}
	}
	return tokens
}