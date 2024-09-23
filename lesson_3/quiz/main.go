package main

import (
	"fmt"
	"os"
	"quiz/questions"
	"quiz/shuffler"
	"quiz/game"
)

func main() {
	questions, err := questions.LoadQuestions()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't load questions: %v\n", err)
		os.Exit(1)
	}

	shuffler.Shuffle(questions)

	correctAnswers := game.Run(questions)

	fmt.Printf("Correct answers: %d", correctAnswers)
}
