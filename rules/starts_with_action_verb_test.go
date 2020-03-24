package rules

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddIsValidActionVerb(t *testing.T) {
	assert.True(t, StartsWithActionVerb("Add something to this new repository and test other things"))
}

func TestIgnoresSpecialCharacters(t *testing.T) {
	assert.True(t, StartsWithActionVerb("[123] Add else"))
	assert.True(t, StartsWithActionVerb("RR-123 Add else"))
}

func TestNonValidVerb(t *testing.T) {
	assert.False(t, StartsWithActionVerb("Cooking new stuff yeet!!"))
}
