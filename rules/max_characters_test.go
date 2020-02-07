package rules

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestFunctionReturnsFalseIfMoreThan80Characters(t *testing.T) {
	var tooLongString = strings.Repeat("A", 73)
	assert.False(t, MaxCharacters(tooLongString, 72), "Should return false")
}

func TestFunctionReturnsTrueIfLessOrEqual80Characters(t *testing.T) {
	var tooLongString = strings.Repeat("A", 72)
	assert.True(t, MaxCharacters(tooLongString, 72), "Should return false")
}
