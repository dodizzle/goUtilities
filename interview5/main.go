package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

type fileInfo struct {
	modTime time.Time
	isDir   string
	size    int64
}

func checkDir(dir string) map[string]fileInfo {
	dirListing := make(map[string]fileInfo)
	var fType string
	contents, err := os.ReadDir(dir)
	if err != nil {
		log.Println(err)
	}
	for _, v := range contents {
		info, err := v.Info()
		if err != nil {
			fmt.Println(err)
		}
		if info.IsDir() {
			fType = "directory"
		} else {
			fType = "file"
		}
		dirListing[info.Name()] = fileInfo{modTime: info.ModTime(), isDir: fType, size: info.Size()}
	}
	for k, v := range dirListing {
		fmt.Printf("name:%s || modTime:%v\n", k, v.modTime)
	}
	return dirListing
}

func main() {
	// Write a Go function that monitors a directory for file changes
	//	(e.g., new files, modifications, or deletions) and logs the events to the console.
	//	Use a simple polling approach with a specified interval (e.g., every 2 seconds)
	// and track file states using their names and last modified times."
	dirListing := checkDir("testDir")
	fmt.Println(dirListing)
}
