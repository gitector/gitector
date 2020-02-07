package rules

import (
	"gitlab.com/tbacompany/gitector/reader"
	"strings"
)

type NotValidEmails struct{}

func CheckForValidEmail(email string, domains []string) bool {
	var emailAndDomain = strings.Split(email, "@")
	if len(emailAndDomain) != 2 {
		return false
	}
	if len(domains) == 0 {
		return true
	}

	for i := 0; i < len(domains); i++ {
		if domains[i] == emailAndDomain[1] {
			return true
		}
	}
	return false
}

func CheckForValidEmailError(email string, domains []string, commit reader.GitCommit) GitError {
	return GitError{
		ErrorType:   NotValidEmails{},
		Title:       "E-mail not part of allowed list",
		Description: "Email " + email + "not on allowed list",
		Commit:      commit,
	}
}
