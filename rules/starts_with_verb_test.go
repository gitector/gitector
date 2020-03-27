package rules

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChecksForBaseForm(t *testing.T) {
	assert.True(t, StartsWithVerb("Add new test"))
}

func TestFailsIfPastTenseUsed(t *testing.T) {
	assert.False(t, StartsWithVerb("Added new test"))
}
