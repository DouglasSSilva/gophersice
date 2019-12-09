package game_test

import (
	"fmt"
	"gophersices/quizgame/commons"
	"gophersices/quizgame/game"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuizzQuestions(t *testing.T) {

	tt := []struct {
		Correct int          `json:"correct"`
		P       game.Problem `json:"problem"`
		Answer  string       `json:"answer"`
	}{}
	err := commons.GetJSONTestFiles(&tt, "../game/testfiles/input.json")
	if err != nil {
		t.Fatalf("Not possible to parse JSON: %v", err)
	}
	correct := 0
	for i, tc := range tt {
		fmt.Println(tc)
		// uses os.Stding and a tempfile
		// to mock user input and answer to quizzGame
		answer := []byte(tc.Answer)
		answerFile, err := ioutil.TempFile("", "answer")

		if err != nil {
			t.Fatalf("Not possible to set answer via stdin")
		}

		defer os.Remove((answerFile.Name()))
		defer answerFile.Close()
		if _, err := answerFile.Write(answer); err != nil {
			t.Fatalf("Not possible to write on answer file")
		} else if _, err := answerFile.Seek(0, 0); err != nil {
			t.Fatalf("Information not found on answer file")
		}

		oldStdin := os.Stdin
		defer func() { os.Stdin = oldStdin }()

		os.Stdin = answerFile

		if err := game.QuizzQuestions(tc.P, &correct, i); err == nil {
			assert.Equal(t, tc.Correct, correct)
		}

	}
}
