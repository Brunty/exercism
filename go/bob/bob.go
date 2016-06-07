package bob // package name must match the package name in bob_test.go

import (
//"regexp"
)
import (
	"strings"
	"unicode"
)

const testVersion = 2 // same as targetTestVersion

func Hey(input string) string {
	var greeting string = strings.TrimSpace(input)

	switch {
	case greeting == "":
		return "Fine. Be that way!"
	case inputIs(greeting, unicode.IsUpper) && !inputIs(greeting, unicode.IsLower):
		return "Whoa, chill out!"
	case greeting[len(greeting)-1] == '?':
		return "Sure."
	default:
		return "Whatever."
	}
}

func inputIs(input string, test func(rune) bool) bool {

	for _, items := range input {
		if test(items) {
			return true
		}
	}

	return false
}
