package rules

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddIsValidActionVerb(t *testing.T) {
	acceptedVerbs := []string{"update", "Add"}
	assert.True(t, StartsWithActionVerb("Add something to this new repository and test other things", acceptedVerbs))
}

func TestIgnoresSpecialCharacters(t *testing.T) {
	acceptedVerbs := []string{"update", "Add"}
	assert.True(t, StartsWithActionVerb("[123] Add else", acceptedVerbs))
	assert.True(t, StartsWithActionVerb("RR-123 Add else", acceptedVerbs))
}

func TestNonValidVerb(t *testing.T) {
	acceptedVerbs := []string{"update", "Add"}
	assert.False(t, StartsWithActionVerb("Cooking new stuff yeet!!", acceptedVerbs))
}
