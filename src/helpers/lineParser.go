package helpers

import "strings"

func ParseQuiz(lines [][]string) []Problem {
	var quiz = make([]Problem, len(lines))

	for i, line := range lines {
		quiz[i] = Problem{
			Question: line[0],
			Answer:   strings.TrimSpace(line[1]),
		}
	}
	return quiz
}

type Problem struct {
	Question string
	Answer   string
}
