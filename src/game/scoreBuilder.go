package game

import "fmt"

func BuildScore(numberOfCorrect, numberOfQuestions int) string {
	return fmt.Sprintf("You scored %d out of %d.\n", numberOfCorrect, numberOfQuestions)
}

func BuildScorePercentage(numberOfCorrect, numberOfQuestions int) int {
	var percentageCorrect = (numberOfCorrect * 100) / numberOfQuestions
	return percentageCorrect
}