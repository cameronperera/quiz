package tests

import (
	"testing"

	"github.com/cameronperera/quiz/src/game"
	"github.com/stretchr/testify/assert"
)

func TestBuildScore(t *testing.T) {
	var expected = "You scored 1 out of 2.\n"
	assert.Equal(t, expected, game.BuildScore(1, 2))
}
