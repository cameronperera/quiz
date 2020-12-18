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

func TestBuildScorePercentage(t *testing.T) {
	assert.Equal(t, 20, game.BuildScorePercentage(2, 10))
	assert.Equal(t, 16, game.BuildScorePercentage(2, 12))
	assert.Equal(t, 23, game.BuildScorePercentage(3, 13))
	assert.Equal(t, 25, game.BuildScorePercentage(1, 4))
}
