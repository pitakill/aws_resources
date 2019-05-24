package main

import (
	"strings"
	"unicode"
)

func containsString(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}

	return false
}

func getKind(s string) string {
	s = strings.TrimFunc(s, func(r rune) bool {
		return unicode.IsLetter(r)
	})

	s = strings.Trim(s, ":")

	s = strings.ToLower(s)

	return s
}
