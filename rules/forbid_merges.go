package rules

import "gitlab.com/tbacompany/gitector/reader"

type ContainsMergeCommit struct{}

func ForbidMerges(isMerge bool, isEnabled bool) bool {
	return isMerge && isEnabled
}

func ContainsMergeCommitError(commit reader.GitCommit) GitError {
	return GitError{
		ErrorType:   ContainsMergeCommit{},
		Title:       "Commit is a merge commit",
		Description: "Merge commits are disabled please use rebase instead",
		Commit:      commit,
	}
}
