package rules

import (
	"gitlab.com/tbacompany/gitector/reader"
	"regexp"
)

type NoTicketNumber struct{}

func ContainsTicketNumber(title string, pattern string) bool {
	var matched, _ = regexp.MatchString(pattern, title)
	return matched
}

func ContainsTicketNumberError(title string, pattern string, commit reader.GitCommit) GitError {
	return GitError{
		ErrorType:   NoTicketNumber{},
		Title:       "Commit doesn't contain ticket number",
		Description: "Couldn't find ticket with pattern " + pattern,
		Commit:      commit,
	}
}
