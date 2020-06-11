package printer

import (
	"encoding/json"
	"gitlab.com/tbacompany/gitector/rules"
)

type jsonWrapper struct {
	Errors []jsonError `json:"errors"`
}

type jsonError struct {
	Error       string `json:"error"`
	CommitTitle string `json:"commit_title"`
}

func GenerateJSON(errors []rules.GitError) string {
	jsonErrors := []jsonError{}

	for _, element := range errors {
		jsonErrors = append(jsonErrors, jsonError{
			Error:       element.Title,
			CommitTitle: element.Commit.Title,
		})

	}
	bytes, _ := json.Marshal(jsonWrapper{jsonErrors})
	return string(bytes)
}
