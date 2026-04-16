package helpers

import(
	"regexp"
	"strings"
)

func ApplyQuote(tokens []string) []string {
	token := strings.Join(tokens, " ")

	token = regexp.MustCompile(`'\s*(.*?)\s*'`).ReplaceAllString(token, " '$1' ")

	token = regexp.MustCompile(`"\s*(.*?)\s*"`).ReplaceAllString(token, "\"$1\" ")

	return strings.Fields(token)
}