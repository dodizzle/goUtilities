package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

// packages to use are flags, csv,os, channels, go routines, time
func getArgs() (*string, *int) {
	fName := flag.String("file", "problems.csv", "a csv file in the format of 'question,answer'")
	timeLimit := flag.Int("limit", 10, "the time limit for the quiz in seconds")
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

func main() {
	fileName, _ := getArgs()
	quiz := map[string]string{}
	parseCSV(*fileName, quiz)
	var userResp string
	var scoreCard []bool
	for question, answer := range quiz {
		fmt.Printf("Question: %s\n", question)
		fmt.Scan(&userResp)
		result := testAnswer(userResp, answer)
		scoreCard = append(scoreCard, result)
	}
	correct := 0
	incorrect := 0
	for _, v := range scoreCard {
		if v {
			correct++
		} else {
			incorrect++
		}
	}
	fmt.Printf("Final score: %d correct, %d incorrect\n", correct, incorrect)
}
