package templates

import (
	"regexp"
	"strings"
	"unicode"
)

// FirstLower sets the first character of the string to lower case
func FirstLower(str string) string {
	if str == "" {
		return ""
	}

	return strings.ToLower(str[0:1]) + str[1:]
}

// FirstUpper sets the first character of the string to upper case
func FirstUpper(str string) string {
	if str == "" {
		return ""
	}
	return strings.ToUpper(str[0:1]) + str[1:]
}

// ToPascalCase returns a pascal-cased (e.g. SomeValueLikeThis) out of a string
func ToPascalCase(value string) string {
	b := strings.Builder{}

	var toUpper bool

	for i, rune := range value {
		// Always upper the first character
		if i == 0 {
			toUpper = true
		}
		// Always upper the character after non-letter/non-digit skipping the character
		if !unicode.IsLetter(rune) && !unicode.IsDigit(rune) {
			toUpper = true
			continue
		}
		// If the flag was set by one of the previous steps
		if toUpper {
			rune = unicode.ToUpper(rune)
			toUpper = false
		}

		b.WriteRune(rune)
	}

	return b.String()
}

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

// ToSnakeCase converts the
func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}
