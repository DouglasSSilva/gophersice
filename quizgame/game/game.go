package game

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

// Problem based on the csv has a question and answer fields/
// to be compared with the user input
type Problem struct {
	Question string
	Answer   string
}

// ParseLines transform a matrix of strings into
// an array of Problem struct;
func ParseLines(lines [][]string) []Problem {
	ret := make([]Problem, len(lines))
	for i, line := range lines {
		ret[i] = Problem{
			Question: line[0],
			Answer:   strings.TrimSpace(line[1]),
		}
	}
	return ret
}

//QuizzQuestions implement a check wether the answer
// provided from the user is the answer for the defined problem.
func QuizzQuestions(p Problem, correct *int, i int) error {
	//uses a scanner to get information from the user
	// will be replaced by an __io.Reader__
	fmt.Printf("Problem #%d: %s = \n", i, p.Question)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	answer := scanner.Text()

	if answer == p.Answer {
		*correct++
	}
	return nil
}

//GetCsv based  on the fileType and name return a csv file
func GetCsv(fileType, fileName string) *os.File {
	csvFileName := flag.String(fileType, fileName, "A csv file in the format of 'question, answer'")

	flag.Parse()
	file, err := os.Open(*csvFileName)

	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFileName))
	}
	return file
}

//QuizzGame a game which receives a csv with problems and expect the user to
// answer correctly all on the terminal.
func QuizzGame() {

	file := GetCsv("csv", "problems.csv")
	r := csv.NewReader(file)
	lines, err := r.ReadAll()

	if err != nil {
		exit("Failed to parse the provided CSV file.")
	}

	problems := ParseLines(lines)
	correct := 0
	for i, p := range problems {
		QuizzQuestions(p, &correct, i)
	}

	fmt.Printf("You scored %d out of %d \n", correct, len(problems))
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
