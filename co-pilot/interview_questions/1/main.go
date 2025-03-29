package main

// Write a Go program that reads a text file and counts
// the number of occurrences of each word in the file.
// The program should:
// Accept the file name as a command-line argument.
// Print each word and its count in alphabetical order.

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/fatih/color"
	"github.com/rodaine/table"
)

type Options struct {
	fileName *string
}

func check(err error) {
	fmt.Println(err)
	os.Exit(1)
}

func cleanStrings(line string) []string {
	line = strings.ToLower(line)
	words := strings.Fields(line)
	return words
}

func main() {
	opts := Options{}
	opts.fileName = flag.String("file", "", "Path to the text file")
	flag.Parse()

	file, err := os.Open(*opts.fileName)
	if err != nil {
		check(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	total := make(map[string]int)
	for scanner.Scan() {
		line := scanner.Text()
		cleanWords := cleanStrings(line)
		for _, v := range cleanWords {
			_, ok := total[v]
			if ok {
				sum := total[v]
				sum++
				total[v] = sum
			} else {
				total[v] = 1
			}
		}
	}
	sortedKeys := make([]string, 0, len(total))
	for k := range total {
		sortedKeys = append(sortedKeys, k)
	}

	sort.Strings(sortedKeys)
	// for _, v := range sortedKeys {
	// 	fmt.Printf("Word: %s || Count: %d\n", v, total[v])
	// }
	tbl := table.New("Word", "Count")
	headerFmt := color.New(color.Underline, color.Bold).SprintfFunc()
	tbl.WithHeaderFormatter(headerFmt)
	for _, v := range sortedKeys {
		tbl.AddRow(v, total[v])
	}
	tbl.Print()
}
