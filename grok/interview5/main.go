package main

import (
	"fmt"
	"log"
	"maps"
	"os"
	"time"
)

func directoryContents(dir string) map[string]time.Time {
	dirListing := make(map[string]time.Time)
	contents, err := os.ReadDir(dir)
	if err != nil {
		log.Println(err)
	}
	for _, v := range contents {
		info, err := v.Info()
		if err != nil {
			fmt.Println(err)
		}
		dirListing[info.Name()] = info.ModTime()
	}
	return dirListing
}

func monitorDir(delay time.Duration) {
	time.Sleep(delay * time.Second)
}

func lookForDeleted(dirListing1 map[string]time.Time, dirListing2 map[string]time.Time) {
	for k, value := range dirListing1 {
		_, ok := dirListing2[k]
		if !ok {
			log.Printf("File deleted %s (%v)\n", k, value)
		}
	}
}

func lookForNew(dirListing1 map[string]time.Time, dirListing2 map[string]time.Time) {
	for k, v := range dirListing2 {
		_, ok := dirListing1[k]
		if !ok {
			log.Printf("File created %s (%v)\n", k, v)
		}
	}
}

func main() {
	// Write a Go function that monitors a directory for file changes
	//	(e.g., new files, modifications, or deletions) and logs the events to the console.
	//	Use a simple polling approach with a specified interval (e.g., every 2 seconds)
	// and track file states using their names and last modified times."
	dir := "testDir"
	pollingInterval := 2
	log.Printf("Checking %s every %d seconds", dir, pollingInterval)
	for {
		dirListing1 := directoryContents(dir)
		monitorDir(2)
		dirListing2 := directoryContents(dir)
		isEqual := maps.Equal(dirListing1, dirListing2)
		if !isEqual {
			lookForDeleted(dirListing1, dirListing2)
			lookForNew(dirListing1, dirListing2)
		}
	}
}
