package reader

import (
	. "gopkg.in/src-d/go-git.v4"
	"io/ioutil"
	"os/exec"
)

func ReadDirectInput(directory string) []GitCommit {
	r, _ := PlainOpen(directory)
	w, _ := r.Worktree()
	s, _ := w.Status()
	changedFiles := 0

	for _, v := range s {
		if v.Staging == Unmodified || v.Staging == Untracked {
			continue
		}
		changedFiles += 1
	}

	out, _ := exec.Command("git", "var", "GIT_AUTHOR_IDENT").Output()
	rawAuthor := string(out)
	commit, _ := ioutil.ReadFile(".git/COMMIT_EDITMSG")
	rawCommit := string(commit)
	return []GitCommit{StringToModel(rawCommit, GitVarToSignature(rawAuthor), changedFiles, false)}
}
