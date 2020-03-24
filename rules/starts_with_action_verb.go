package rules

import (
	"gitlab.com/tbacompany/gitector/reader"
	"regexp"
	"strings"
)

type NoActionVerb struct{}

func extractFirstWord(title string) string {
	r := regexp.MustCompile("[a-zA-Z]+")
	return r.FindString(title)
}

func contains(s string, arr []string) bool {
	for _, word := range arr {
		if word == s {
			return true
		}
	}
	return false
}

func StartsWithActionVerb(title string) bool {
	ACTION_VERBS := []string{"add", "remove", "update", "fix", "move", "delete"}

	firstWord := extractFirstWord(title)
	return contains(strings.ToLower(firstWord), ACTION_VERBS)
}

func StartsWithActionVerbError(title string, commit reader.GitCommit) GitError {
	return GitError{
		ErrorType:   NoActionVerb{},
		Title:       "Commit doesn't start with action verb",
		Description: "Commit should start with action verb",
		Commit:      commit,
	}
}
