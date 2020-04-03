package reader

import (
	"golang.org/x/net/http2"
	. "gopkg.in/src-d/go-git.v4"
	. "gopkg.in/src-d/go-git.v4/_examples"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"strings"
)

func ReadGitCommitsFromDirectory(directory string, scope string) []GitCommit {
	r, err := PlainOpen(directory)
	CheckIfError(err)

	if scopeBetweenBranches(scope) {
		a, b := extractBranches(scope)

		var models []GitCommit
		iterateOverCommon(r, getHashFromBranchName(r, a), getHashFromBranchName(r, b), func(c *object.Commit) error {
			models = append(models, StringToModel(c.Message, Signature{Email: c.Author.Email, Name: c.Author.Name}, countFilesInCommit(c)))
			return nil
		})
		return models
	} else {
		return []GitCommit{}
	}

}

func scopeBetweenBranches(scope string) bool {
	return strings.Contains(scope, "..")
}

func extractBranches(scope string) (string, string) {
	return strings.Split(scope, "..")[0], strings.Split(scope, "..")[1]
}

func getHashFromBranchName(r *Repository, branchName string) plumbing.Hash {
	if branchName == "" {
		c, _ := r.Head()
		return c.Hash()
	} else {
		var ref plumbing.Hash
		iter, _ := r.References()
		iter.ForEach(func(reference *plumbing.Reference) error {
			if strings.HasSuffix(reference.Name().String(), branchName) {
				ref = reference.Hash()
			}
			return nil
		})
		return ref
	}
}

func iterateOverCommon(r *Repository, compared plumbing.Hash, comparing plumbing.Hash, cb func(*object.Commit) error) {
	cIter, _ := r.Log(&LogOptions{From: comparing})

	comparingSameBranches := comparing == compared

	cIter.ForEach(func(c *object.Commit) error {
		if containsCommonCommit(r, c, compared) && !comparingSameBranches {
			return http2.GoAwayError{}
		} else {
			return cb(c)
		}
	})
}

func containsCommonCommit(r *Repository, commit *object.Commit, compared plumbing.Hash) bool {
	cIter, _ := r.Log(&LogOptions{From: compared})
	res := false
	cIter.ForEach(func(c *object.Commit) error {
		if c.Hash == commit.Hash {
			res = true
		}
		return nil
	})
	return res
}

func countFilesInCommit(commit *object.Commit) int {
	count := 0

	files, _ := commit.Files()
	files.ForEach(func(file *object.File) error {
		count += 1
		return nil
	})

	return count
}
