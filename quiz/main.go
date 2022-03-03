package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

type Problem struct {
	question string
	answer   string
}

func main() {
	filepath := flag.String("csv", "problems.csv", "a csv file in format of 'question,answer'")
	flag.Parse()

	problems := readProblemsFrom(*filepath)

	score := 0

	for i, problem := range problems {
		fmt.Printf("Problem #%d - %s: ", i+1, problem.question)
		var ans string
		fmt.Scanf("%s\n", &ans)

		if ans == problem.answer {
			score++
		}
	}

	fmt.Print("\n\n")
	fmt.Printf("You scored %v out of %v\n", score, len(problems))
}

func readProblemsFrom(p string) []Problem {

	problems := []Problem{}

	f, err := os.Open(p)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s", p))
	}
	defer f.Close()

	r := csv.NewReader(f)

	records, err := r.ReadAll()

	if err != nil {
		exit(fmt.Sprintf("Failed to parse the CSV file: %s", p))
	}

	for _, record := range records {
		problems = append(problems, Problem{
			question: record[0],
			answer:   strings.TrimSpace(record[1]),
		})
	}

	return problems
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
