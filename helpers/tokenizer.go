package helpers

import(
	"regexp"
)

var t = regexp.MustCompile(`\([^)]+\)|\S+`)

func TokenizeElements(tokens  string) []string {
	if len(tokens) == 0 {
		return []string{}
	}
	token := t.FindAllString(tokens, -1)
	return token
}