package rules

import (
	"gitlab.com/tbacompany/gitector/reader"
	"strings"
)

type TrailingPunctationInTitle struct{}

func CheckForTrailingPunctuationInTitle(title string) bool {
	parts := strings.Split(title, ".")
	if len(parts) == 1 {
		return true
	}
	if strings.TrimSpace(parts[len(parts)-1]) == "" {
		return false
	}
	return true
}

func CheckForTrailingPunctuationInTitleError(commit reader.GitCommit) GitError {
	return GitError{
		ErrorType:   NotValidEmails{},
		Title:       "Trailing Punctuation not allowed in title",
		Description: "Commit title should not contain trailing punctuation",
		Commit:      commit,
	}
}
