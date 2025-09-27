package main

import (
	"fmt"
	"os"
	"quiz/game"
	"quiz/questions"
	"quiz/shuffler"
)

const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
)

func main() {
	questions, err := questions.LoadQuestions()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't load questions %v\n", err)
		os.Exit(1)
	}

	shuffler.Shuffle(questions)
	correctAnswers, incorrectAnswers, percent := game.Run(questions)

	fmt.Println("Results:")
	fmt.Println("correctAnswers: ", ColorGreen, correctAnswers, ColorReset)
	fmt.Println("incorrectAnswers: ", ColorRed, incorrectAnswers, ColorReset)

	var percentColor string
	if percent >= 70 {
		percentColor = ColorGreen
	} else if percent >= 50 {
		percentColor = ColorYellow
	} else {
		percentColor = ColorRed
	}
	fmt.Printf("Success rate: %s%.4f%%%s\n", percentColor, percent, ColorReset)
}
