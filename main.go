package main

import (
	"flag"
	"fmt"
	"encoding/csv"
	"os"
	"strings"
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

func main() {
	file := readArgs()
	readCSV(file)
}