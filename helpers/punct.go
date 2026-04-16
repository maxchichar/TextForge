package helpers

import(
	"regexp"
	"strings"
)

func ApplyPunct(tokens []string) []string {
	token := strings.Join(tokens, " ")

	token = regexp.MustCompile(`\s+([.,:;!?])`).ReplaceAllString(token, "$1")
	token = regexp.MustCompile(`([.,:;!?])(\s*)(\w)`).ReplaceAllString(token, "$1 $3")
	return strings.Fields(token)
}