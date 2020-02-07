package printer

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/tbacompany/gitector/rules"
)

func TestConsoleTextContainsPassedErrors(t *testing.T) {
	errors := []rules.GitError{{Title: "test"}, {Title: "test2"}}
	expected := "==== Found errors: ====\n" +
		"\x1b[0;31mtest\x1b[0m: \n" +
		"\x1b[0;31mtest2\x1b[0m: \n" +
		"==== Summary: ====\n" +
		"\x1b[0;31m2 errors found \n\x1b[0m"
	assert.Equal(t, expected, GenerateConsoleErrors(errors))
}
