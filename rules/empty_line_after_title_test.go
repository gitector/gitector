package rules

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSuccessIfEmptyDescription(t *testing.T) {
	description := ""
	assert.True(t, EmptyLineBetween(description))
}

func TestFailsIfDoesntStarWithNewLine(t *testing.T) {
	description := "TestQEWQEwq\nasdsadas"
	assert.False(t, EmptyLineBetween(description))
}

func TestSuccessIfStarsWithNewLine(t *testing.T) {
	description := "\nTestQEWQEwq\nasdsadas"
	assert.True(t, EmptyLineBetween(description))
}
