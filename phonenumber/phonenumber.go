package phonenumber

import (
	"sort"
	"strings"
)

var keyboard = map[string]string{
	"a": "2",
	"b": "2",
	"c": "2",

	"d": "3",
	"e": "3",
	"f": "3",

	"g": "4",
	"h": "4",
	"i": "4",

	"j": "5",
	"k": "5",
	"l": "5",

	"m": "6",
	"n": "6",
	"o": "6",

	"p": "7",
	"q": "7",
	"r": "7",
	"s": "7",

	"t": "8",
	"u": "8",
	"v": "8",

	"w": "9",
	"x": "9",
	"y": "9",
	"z": "9",
}

func convertWordToNumber(rawWord string) string {
	var wordAsNumber string

	for _, c := range rawWord {
		ke := keyboard[string(c)]
		wordAsNumber += ke
	}

	return wordAsNumber
}

func FindMatches(phoneNumber string, words []string) []string {
	var matches []string

	for _, word := range words {
		if strings.Contains(phoneNumber, convertWordToNumber(word)) {
			matches = append(matches, word)
		}
	}

	sort.Strings(matches)

	return matches
}
