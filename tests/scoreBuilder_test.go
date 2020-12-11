package tests

import (
	"testing"

	"github.com/cameronperera/quiz/src/helpers"
	"github.com/stretchr/testify/assert"
)

func TestBuildScore(t *testing.T) {
	var expected = "You scored 1 out of 2.\n"
	assert.Equal(t, expected, helpers.BuildScore(1, 2))
}
