package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type Problem struct {
	question string
	answer   string
}

func main() {
	filepath := flag.String("csv", "problems.csv", "a csv file in format of 'question,answer'")
	duration := flag.Int64("limit", 30, "timeout for every question in seconds")
	flag.Parse()

	problems := readProblemsFrom(*filepath)

	timer := time.NewTimer(time.Duration(*duration) * time.Second)

	score := 0
	for i, problem := range problems {
		fmt.Printf("Problem #%d - %s: ", i+1, problem.question)
		answerCh := make(chan string)
		go func() {
			var ans string
			fmt.Scanf("%s\n", &ans)
			answerCh <- ans
		}()

		select {
		case <-timer.C:
			showScoreAndExit(score, len(problems))
		case ans := <-answerCh:
			if ans == problem.answer {
				score++
			}
		}
	}

	showScoreAndExit(score, len(problems))
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

func showScoreAndExit(score int, total int) {
	fmt.Print("\n\n")
	fmt.Printf("You scored %d out of %d\n", score, total)
	os.Exit(0)
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
