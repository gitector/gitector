package rules

import "gitlab.com/tbacompany/gitector/reader"

type TooLong struct{}

func MaxCharacters(title string, maxLength int) bool {
	return len(title) <= maxLength
}

func MaxCharactersError(title string, maxLength int, commit reader.GitCommit) GitError {
	return GitError{
		ErrorType:   TooLong{},
		Title:       "Commit Title is too long",
		Description: "Commit have more than " + string(maxLength),
		Commit:      commit,
	}
}
