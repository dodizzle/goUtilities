package main

import (
	"fmt"
	"time"
)

func hello(done chan bool) {
	start := 1
	for {
		select {
		case <-done:
			return
		default:
			fmt.Printf("%d:) hello\n", start)
			start++
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	done := make(chan bool)
	go hello(done)
	time.Sleep(5 * time.Second)
	done <- true
	fmt.Println("Exit loop")
}
