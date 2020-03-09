package rules

import (
	"gitlab.com/tbacompany/gitector/reader"
	"strconv"
)

type TooManyFilesInSingleCommit struct{}

func CheckForMaxFiles(filesInCommit int, maxFiles int) bool {
	return filesInCommit <= maxFiles
}

func CheckForMaxFilesError(filesInCommit int, maxFiles int, commit reader.GitCommit) GitError {
	return GitError{
		ErrorType:   TooManyFilesInSingleCommit{},
		Title:       "Too many files in single commit",
		Description: "Changed " + strconv.Itoa(filesInCommit) + " files, max allowed: " + strconv.Itoa(maxFiles),
		Commit:      commit,
	}
}
