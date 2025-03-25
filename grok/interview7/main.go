package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func searchLogs(logFile string, keyWord string) int {
	f, err := os.Open(logFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()
	count := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), keyWord) {
			count++
		}
	}
	return count
}

func main() {
	//Write a Go function that reads a log file and counts the number of lines containing a specific keyword (e.g., 'ERROR').
	//  The function should take the file path and keyword as inputs and return the count.
	// Handle potential errors like a missing file or read issues.
	logFile := os.Args[1]
	if len(logFile) == 0 {
		fmt.Println(" you must supply a log file")
		os.Exit(1)
	}
	keyWord := "ERROR"
	count := searchLogs(logFile, keyWord)
	fmt.Printf("Found %d %ss in %s\n", count, keyWord, logFile)
}
