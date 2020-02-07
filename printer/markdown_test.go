package printer

import (
	"github.com/stretchr/testify/assert"
	"gitlab.com/tbacompany/gitector/rules"
	"testing"
)

func TestMarkdownContainsPassedErrors(t *testing.T) {
	errors := []rules.GitError{{Title: "test"}, {Title: "test2"}}
	expectedRes := `## Errors found: 2

- test
- test2
`
	assert.Equal(t, GenerateMarkdown(errors), expectedRes)
}
