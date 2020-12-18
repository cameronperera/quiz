package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/cameronperera/quiz/src/game"
	"github.com/cameronperera/quiz/src/helpers"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "as csv file in the format of 'question,answer'")
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")
	flag.Parse()

	var lines, err = helpers.ReadFile(*csvFilename)
	if err != nil {
		exit(err.Error())
	}

	problems := helpers.ParseQuiz(lines)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	correct := 0
problemloop:
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, p.Question)
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Println()
			break problemloop
		case answer := <-answerCh:
			if answer == p.Answer {
				correct++
			}
		}
	}

	fmt.Printf(game.BuildScore(correct, len(problems)))
	fmt.Printf("Score: %d%% correct\n", game.BuildScorePercentage(correct, len(problems)))
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
