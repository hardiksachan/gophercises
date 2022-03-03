package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type Problem struct {
	question string
	answer   string
}

func main() {
	filepath := "problems.csv"
	problems, err := readProblemsFrom(filepath)

	if err != nil {
		log.Fatal(err)
	}

	score := 0

	for _, problem := range problems {
		fmt.Printf("%s: ", problem.question)
		var ans string
		fmt.Scan(&ans)

		if ans == problem.answer {
			score++
		}
	}

	fmt.Print("\n\n")
	fmt.Printf("You scored %v out of %v\n", score, len(problems))
}

func readProblemsFrom(p string) ([]Problem, error) {

	problems := []Problem{}

	f, err := os.Open(p)

	if err != nil {
		return problems, err
	}

	defer f.Close()

	r := csv.NewReader(f)

	records, err := r.ReadAll()

	if err != nil {
		return problems, err
	}

	for _, record := range records {
		problems = append(problems, Problem{
			question: record[0],
			answer:   record[1],
		})
	}

	return problems, nil
}
