package main

import (
	"flag"
	"fmt"
	"encoding/csv"
	"os"
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

func readCSV(filename string) {
	file, err := os.Open(filename)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)

	problems, err := reader.ReadAll()

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var quiz []Question

	for _, row := range problems {
		if len(row) < 2 {
			fmt.Println("Skipping malformed row:", row)
			continue
		}

		q := Question{
			question: row[0],
			answer:   row[1],
		}
		quiz = append(quiz, q)
	}

	for _, question := range quiz {
		fmt.Printf("Question: %s, Answer: %s\n", question.question, question.answer)
	}
}

func main() {
	file := readArgs()
	readCSV(file)
}