package main

import (
	"flag"
	"fmt"
	"encoding/csv"
	"os"
	"strings"
	"bufio"
)

type Question struct {
	question string
	answer string
}


func readArgs() string {
	filename := flag.String("filename", "problems.csv", "a csv file in the format of 'question,answer'")
	
	flag.Parse()

	return *filename

}

func readCSV(filename string) ([]Question, error) {
	file, err := os.Open(filename)

	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	problems, err := reader.ReadAll()

	if err != nil {
		return nil, fmt.Errorf("failed to read CSV data: %w", err)
	}

	var quiz []Question

	for _, row := range problems {
		if len(row) < 2 {
			fmt.Println("Skipping malformed row:", row)
			continue
		}

		q := Question{
			question: strings.TrimSpace(row[0]),
			answer:   strings.TrimSpace(row[1]),
		}
		quiz = append(quiz, q)
	}

	return quiz, nil
}

func startQuiz(quiz []Question) {
	correctAnswers := 0
	reader := bufio.NewReader(os.Stdin)

	for i, question := range quiz {
		fmt.Printf("Question %d: %s = ", i+1, question.question)

		// Get the user's answer
		userAnswer, _ := reader.ReadString('\n')
		userAnswer = strings.TrimSpace(userAnswer) // Remove any extra whitespace

		// Compare answers (case-insensitive)
		if strings.EqualFold(userAnswer, question.answer) {
			fmt.Println("Correct!")
			correctAnswers++
		} else {
			fmt.Printf("Incorrect. The correct answer is %s.\n", question.answer)
		}
	}

	// Display the total score
	fmt.Printf("\nYou scored %d out of %d.\n", correctAnswers, len(quiz))
}

func main() {
	file := readArgs()
	quiz, _ := readCSV(file)
	startQuiz(quiz)
}