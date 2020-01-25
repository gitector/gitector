package rules

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckPassesForPatternInTitle(t *testing.T) {
	assert.True(t, ContainsTicketNumber("[123] This is simple text", "\\[[0-9]+\\]"))
}

func TestCheckFailesForNoPatternInTitle(t *testing.T) {
	assert.False(t, ContainsTicketNumber("<123> This is simple text", "\\[[0-9]+\\]"))
}
