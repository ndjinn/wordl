package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrInList(t *testing.T) {
	list := []string{"able", "baker", "charlie"}
	pos := "able"
	neg := "delta"
	assert.True(t, StrInList(pos, list))
	assert.False(t, StrInList(neg, list))
}
