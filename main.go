package main

import (
	"flag"
	"fmt"
	"encoding/csv"
	"os"
)


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

	for _, row := range problems {
		if len(row) < 2 {
			fmt.Println("Skipping malformed row:", row)
			continue
		}
		question := row[0]
		answer := row[1]
		fmt.Printf("Question: %s, Answer: %s\n", question, answer)
	}
}

func main() {
	file := readArgs()
	readCSV(file)
}