package reader

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestSplitsCorrectlyTitleFromMessage(t *testing.T) {
	var message = `this is simple commit title

some body bla bla bla 
body 
body

body continues
and continues




and continues
`
	commit := StringToModel(message, Signature{})
	assert.Equal(t, len(strings.Split(commit.Title, "\n")), 1)
	assert.Equal(t, len(strings.Split(commit.Description, "\n")), 13)
}

func TestSplitsCorrectlyTitleFromMessageWhenNoBreak(t *testing.T) {
	var message = `this is simple commit title
some body bla bla bla 
body 
body
body continues
and continues
and continues
`
	commit := StringToModel(message, Signature{})
	assert.Equal(t, len(strings.Split(commit.Title, "\n")), 1)
	assert.Equal(t, len(strings.Split(commit.Description, "\n")), 7)
}

func TestParsingEmailFromGitVar(t *testing.T) {
	var gitVar = `James Bond <james@mi7.co.uk> 1575497822 +0100`
	assert.Equal(t, "James Bond", GitVarToSignature(gitVar).Name)
	assert.Equal(t, "james@mi7.co.uk", GitVarToSignature(gitVar).Email)
}
