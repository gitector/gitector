package printer

import (
	"github.com/stretchr/testify/assert"
	"gitlab.com/tbacompany/gitector/reader"
	"gitlab.com/tbacompany/gitector/rules"
	"testing"
)

func TestJsonContainesPassedErrors(t *testing.T) {
	errors := []rules.GitError{{Title: "test", Commit: reader.GitCommit{Title: "Test"}}, {Title: "test2", Commit: reader.GitCommit{Title: "Test2"}}}
	expectedRes := `{"errors":[{"error":"test","commit_title":"Test"},{"error":"test2","commit_title":"Test2"}]}`
	assert.Equal(t, expectedRes, GenerateJSON(errors))
}
