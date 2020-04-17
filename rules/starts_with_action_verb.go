package rules

import (
	"gitlab.com/tbacompany/gitector/reader"
	"gitlab.com/tbacompany/gitector/utils"
)

type NoActionVerb struct{}

func StartsWithActionVerb(title string, acceptedActionVerbs []string) bool {
	firstWord := utils.ExtractFirstWord(title)
	return utils.Contains(firstWord, acceptedActionVerbs)
}

func StartsWithActionVerbError(title string, commit reader.GitCommit) GitError {
	return GitError{
		ErrorType:   NoActionVerb{},
		Title:       "Commit doesn't start with action verb",
		Description: "Commit should start with action verb",
		Commit:      commit,
	}
}
