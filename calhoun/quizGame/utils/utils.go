package utils

import (
	"log"
	"os"
)

func OpenCsv(fileName string) {
	file, err := os.ReadFile(fileName) // For read access.
	if err != nil {
		log.Fatal(err)
	}
	os.Stdout.Write(file)
}
