package reader

import (
	"io/ioutil"
	"os/exec"
)

func ReadDirectInput() []GitCommit {
	out, _ := exec.Command("git", "var", "GIT_AUTHOR_IDENT").Output()
	rawAuthor := string(out)
	commit, _ := ioutil.ReadFile(".git/COMMIT_EDITMSG")
	rawCommit := string(commit)
	return []GitCommit{StringToModel(rawCommit, GitVarToSignature(rawAuthor))}
}
