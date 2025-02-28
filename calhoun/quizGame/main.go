package main

import (
	"fmt"

	"github.com/dodizzle/goUtilities/calhoun/quizGame/utils"
)

const (
	csvFileName = "quiz.csv"
)

func main() {
	fmt.Println("Hello")
	utils.OpenCsv(csvFileName)

}
