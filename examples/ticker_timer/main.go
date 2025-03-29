package main

import (
	"fmt"
	"time"
)

func run(ticker *time.Ticker, timer *time.Timer) {
	start := 1
	for {
		select {
		case <-ticker.C:
			fmt.Printf("%v %v\n", start, time.Now())
			start++
		case <-timer.C:
			fmt.Println("Finished!")
			return
		}
	}
}

func main() {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()
	timer := time.NewTimer(10 * time.Second)
	defer timer.Stop()
	fmt.Println("Starting...")
	run(ticker, timer)
}
