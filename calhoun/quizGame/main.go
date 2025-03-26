package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

var userResp string

// packages to use are flags, csv,os, channels, go routines, time
func getArgs() (*string, *time.Duration) {
	fName := flag.String("file", "problems.csv", "a csv file in the format of 'question,answer'")
	timeLimit := flag.Duration("limit", 10*time.Second, "the time limit for the quiz in seconds")
	flag.Parse()
	return fName, timeLimit
}

func parseCSV(fileName string, quiz map[string]string) {
	csvFile, error := os.ReadFile(fileName)
	if error != nil {
		fmt.Printf("Error opening filename: %s\n", fileName)
		os.Exit(1)
	}

	r := csv.NewReader(strings.NewReader(string(csvFile)))
	for {
		record, error := r.Read()
		if error == io.EOF {
			break
		}
		if error != nil {
			fmt.Println("Error reading csv file")
			os.Exit(1)
		}
		quiz[record[0]] = record[1]
	}
}

func testAnswer(r string, a string) bool {
	return r == a
}

func runQuiz(quiz map[string]string, done chan bool) {
	var scoreCard []bool
	correct := 0
	incorrect := 0
	for question, answer := range quiz {
		select {
		case <-done:
			for _, v := range scoreCard {
				if v {
					correct++
				} else {
					incorrect++
				}
			}
			fmt.Printf("Final score: %d correct, %d incorrect\n", correct, incorrect)
			return
		default:
			fmt.Printf("Question: %s\n", question)
			fmt.Scan(&userResp)
			result := testAnswer(userResp, answer)
			scoreCard = append(scoreCard, result)
		}
	}
}

func main() {
	fileName, timeLimit := getArgs()
	quiz := map[string]string{}
	parseCSV(*fileName, quiz)
	done := make(chan bool)
	go runQuiz(quiz, done)
	time.Sleep(*timeLimit)
	done <- true
	time.Sleep(100 * time.Millisecond)
}
