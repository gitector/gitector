package rules

import (
	"gitlab.com/tbacompany/gitector/reader"
	strings "strings"
)

type NoEmptyLine struct{}

func EmptyLineBetween(description string) bool {
	return description == "" || strings.HasPrefix(description, "\n")
}

func EmptyLineBetweenError(commit reader.GitCommit) GitError {
	return GitError{
		ErrorType:   NoEmptyLine{},
		Title:       "Missing empty line after title",
		Description: "If your commit message contains description you should split it with empty line",
		Commit:      commit,
	}
}
