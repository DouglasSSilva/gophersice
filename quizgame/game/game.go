package game

import (
	"bufio"
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

var mux sync.Mutex

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
func QuizzQuestions(p Problem, correct *int, i, totalProblems int, timer time.Timer) error {
	//uses a scanner to get information from the user
	// will be replaced by an __io.Reader__

	fmt.Printf("Problem #%d: %s = \n", i+1, p.Question)
	var answer string
	answerCh := make(chan string)

	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		answer = scanner.Text()
		answerCh <- answer
	}()

	select {
	case <-timer.C:
		fmt.Printf("You scored %d out of %d \n", *correct, totalProblems)
		err := errors.New("tempo esgotado")
		return err
	case answer := <-answerCh:
		if answer == p.Answer {
			*correct++
		}
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

	timeLimit := flag.Int("limit", 30, "The time limit for the quiz in seconds")

	file := GetCsv("csv", "problems.csv")
	r := csv.NewReader(file)
	lines, err := r.ReadAll()

	if err != nil {
		exit("Failed to parse the provided CSV file.")
	}

	problems := ParseLines(lines)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	var correct int
	for i, p := range problems {
		err := QuizzQuestions(p, &correct, i, len(problems), *timer)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
