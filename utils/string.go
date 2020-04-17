package utils

import (
	"regexp"
	"strings"
)

func ExtractFirstWord(title string) string {
	words := strings.Split(title, " ")
	isLetters := regexp.MustCompile(`^[a-zA-Z]+$`).MatchString
	for _, word := range words {
		if isLetters(word) {
			return word
		}
	}
	return ""
}

func Contains(s string, arr []string) bool {
	for _, word := range arr {
		if strings.ToLower(word) == strings.ToLower(s) {
			return true
		}
	}
	return false
}
