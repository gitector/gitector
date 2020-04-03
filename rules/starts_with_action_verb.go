package rules

import (
	"gitlab.com/tbacompany/gitector/reader"
	"regexp"
	"strings"
)

type NoActionVerb struct{}

func extractFirstWord(title string) string {
	words := strings.Split(title, " ")
	isLetters := regexp.MustCompile(`^[a-zA-Z]+$`).MatchString
	for _, word := range words {
		if isLetters(word) {
			return word
		}
	}
	return ""
}

func contains(s string, arr []string) bool {
	for _, word := range arr {
		if strings.ToLower(word) == strings.ToLower(s) {
			return true
		}
	}
	return false
}

func StartsWithActionVerb(title string, acceptedActionVerbs []string) bool {
	firstWord := extractFirstWord(title)
	return contains(firstWord, acceptedActionVerbs)
}

func StartsWithActionVerbError(title string, commit reader.GitCommit) GitError {
	return GitError{
		ErrorType:   NoActionVerb{},
		Title:       "Commit doesn't start with action verb",
		Description: "Commit should start with action verb",
		Commit:      commit,
	}
}
