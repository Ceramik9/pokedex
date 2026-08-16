package repl

import (
	"strings"
)


func cleanInput(text string) []string {
	if len(text) == 0 || text == " " {
		return []string{}
	}
	return strings.Split(text, " ")
}
