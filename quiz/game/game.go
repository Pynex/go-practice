package game

import (
	"bufio"
	"fmt"
	"os"
	"quiz/questions"
	"strings"
)

const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
)

func Run(questions []questions.Question) (correctAnswers, incorrectAnswers uint, procent float32) {
	fmt.Println("Welcome to the Country Quiz!")

	for _, question := range questions {
		if askQuestion(question) {
			correctAnswers++
			fmt.Println(ColorGreen + "Correct!" + ColorReset)
		} else {
			incorrectAnswers++
			fmt.Println(ColorRed + "Incorrect!" + ColorReset)
		}
	}

	totalQuestions := correctAnswers + incorrectAnswers
	if totalQuestions > 0 {
		procent = float32(correctAnswers) / float32(totalQuestions) * 100
	} else {
		procent = 0
	}
	return correctAnswers, incorrectAnswers, procent
}

func askQuestion(question questions.Question) bool {
	fmt.Printf("\nEnter the capital of %s: ", question.Country)

	if getUserInput() == strings.ToLower(question.Capital) {
		return true
	} else {
		return false
	}
}

func getUserInput() string {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Your answer: ")
		result, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("An error occured while reading the entered text. Please try again!")
			continue
		}

		return strings.ToLower(strings.TrimRight(result, "\r\n"))
	}
}
