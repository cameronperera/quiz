package helpers

import "fmt"

func BuildScore(numberOfCorrect, numberOfQuestions int) string {
	return fmt.Sprintf("You scored %d out of %d.\n", numberOfCorrect, numberOfQuestions)
}
