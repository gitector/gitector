package rules

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEmailCanBesplitedByAt(t *testing.T) {
	assert.False(t, CheckForValidEmail("test", []string{}), "Wrong email should return false")
}

func TestReturnTrueIfEmailIsOnList(t *testing.T) {
	domains := []string{"tba.com", "pm.com"}
	assert.True(t, CheckForValidEmail("test@tba.com", domains), "Correct email should return true")
}

func TestReturnFalseIfEmailIsNotOnList(t *testing.T) {
	domains := []string{"tba.com", "pm.com"}
	assert.False(t, CheckForValidEmail("test@yolo.com", domains), "Email shouldn't be found")
}
