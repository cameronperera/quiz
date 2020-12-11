package tests

import (
	"testing"

	"github.com/cameronperera/quiz/src/helpers"
	"github.com/stretchr/testify/assert"
)

func TestParseQuiz(t *testing.T) {
	var expected = []helpers.Problem{
		{Question: "1+1", Answer: "2"},
		{Question: "1+2", Answer: "3"},
	}
	lines := [][]string{{"1+1", "2"}, {"1+2", "3"}}
	var actual = helpers.ParseQuiz(lines)

	assert.Equal(t, expected, actual)
}
