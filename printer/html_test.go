package printer

import (
	"github.com/stretchr/testify/assert"
	"gitlab.com/tbacompany/gitector/rules"
	"testing"
)

func TestHtmlContainsPassedErrors(t *testing.T) {
	errors := []rules.GitError{{Title: "test"}, {Title: "test2"}}
	expectedRes := `<!doctype html>

<html lang="en">
<head>
  <meta charset="utf-8">
  <title>Error report</title>
</head>
<body>
    <h1>Errors found: 2</h1>
    <p>test</p>
    <p>test2</p>
</body>
</html>`
	assert.Equal(t, GenerateHTML(errors), expectedRes)
}
