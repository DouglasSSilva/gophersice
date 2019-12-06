package main

import "flag"

import (
	"encoding/csv"
	"fmt"
	"os"
)

//  a list of problems based on the csv files
type Problem struct {
	Question string
	Answer   string
}

func parseLines(lines [][]string) []Problem {
	ret := make([]Problem, len(lines))
	for i, line := range lines {
		ret[i] = Problem{
			Question: line[0],
			Answer:   line[1],
		}
	}
	return ret
}

func main() {
	csvFileName := flag.String("csv", "problems.csv", "A csv file in the format of 'question, answer'")

	flag.Parse()
	file, err := os.Open(*csvFileName)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFileName))
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()

	if err != nil {
		exit("Failed to parse the provided CSV file.")
	}

	problems := parseLines(lines)

	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.Question)
		var answer string
		fmt.Scanf("%s", &answer)
		if answer == p.Answer {
			fmt.Printf("Correct\n")
		} else {
			fmt.Printf("Incorrect\n")
		}
	}
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
