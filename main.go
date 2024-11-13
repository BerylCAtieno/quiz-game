package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

//add timer
// when time stops, the quiz ends even if user is waiting to enter input
// ask user to press enter to start timer
//add shuffle flag to shuffle the questions if called

type Question struct {
	question string
	answer   string
}

func readArgs() (string, int, bool) {
	filename := flag.String("filename", "problems.csv", "a csv file in the format of 'question,answer'")
	timer := flag.Int("timer", 30, "indicates the amount of time, in seconds, to complete a quiz")
	shuffle := flag.Bool("shuffle", false, "indicates whether the questions should be shuffled before the quiz starts")

	flag.Parse()

	return *filename, *timer, *shuffle

}

func readCSV(filename string) ([]Question, error) {

	fullPath := "quizes/" + filename
	file, err := os.Open(fullPath)

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

func startQuiz(quiz []Question, timeLimit int, shuffle bool) {

	if shuffle {
		rand.Seed(time.Now().UnixNano()) // Seed for randomness
		rand.Shuffle(len(quiz), func(i, j int) {
			quiz[i], quiz[j] = quiz[j], quiz[i]
		})
	}

	fmt.Println("Press Enter to start the quiz...")
	bufio.NewReader(os.Stdin).ReadString('\n')
	fmt.Println("Quiz started!")

	timer := time.NewTimer(time.Duration(timeLimit) * time.Second)
	correctAnswers := 0

	answerCh := make(chan string)

quizLoop:
	for i, question := range quiz {
		fmt.Printf("Question %d: %s = \n", i+1, question.question)

		// Launch a goroutine to capture the user's answer asynchronously
		go func() {
			reader := bufio.NewReader(os.Stdin)
			userAnswer, _ := reader.ReadString('\n')
			answerCh <- strings.TrimSpace(userAnswer) // Send trimmed input to channel
		}()

		// Handle answer or timeout
		select {
		case <-timer.C:
			fmt.Println("\nTime's up!")
			break quizLoop
		case userAnswer := <-answerCh:
			// Check the answer and provide feedback
			if strings.EqualFold(userAnswer, question.answer) {
				fmt.Println("Correct!")
				correctAnswers++
			} else {
				fmt.Printf("Incorrect. The correct answer is %s.\n", question.answer)
			}
		}
	}

	// Display the final score
	fmt.Printf("\nYou scored %d out of %d.\n", correctAnswers, len(quiz))
}

func main() {
	file, time, shuffle := readArgs()
	quiz, _ := readCSV(file)
	startQuiz(quiz, time, shuffle)
}
