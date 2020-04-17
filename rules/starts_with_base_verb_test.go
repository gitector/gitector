package rules

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChecksForBaseForm(t *testing.T) {
	assert.True(t, StartsWithBaseVerb("Add new test"))
}

func TestFailsIfPastTenseUsed(t *testing.T) {
	assert.False(t, StartsWithBaseVerb("Added new test"))
}

func TestFailsForIrregular(t *testing.T) {
	assert.False(t, StartsWithBaseVerb("took new actions against further faliures"))
}
