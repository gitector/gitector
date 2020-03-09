package rules

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSuccessIfPassesForEmptyTitle(t *testing.T) {
	title := ""
	assert.True(t, CheckForTrailingPunctuationInTitle(title))
}

func TestFailsIfTrailingPunctuationInTitle(t *testing.T) {
	title := "Add new feature."
	assert.False(t, CheckForTrailingPunctuationInTitle(title))
}

func TestSuccessIfNoTrailingPunctuation(t *testing.T) {
	title := "Add new feature"
	assert.True(t, CheckForTrailingPunctuationInTitle(title))
}
