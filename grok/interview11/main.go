package main

import (
	"fmt"
	"time"
)

func runTask(done chan bool) {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			fmt.Println("Task running at", time.Now().Format("15:04:05"))
		case <-done:
			fmt.Println("Task stopped")
			return
		}
	}
}

func main() {
	done := make(chan bool) // Channel to signal task completion

	// Start the task in a goroutine
	go runTask(done)

	// Run for 5 seconds
	time.Sleep(5 * time.Second)

	// Signal the task to stop
	done <- true

	// Give it a moment to finish printing
	time.Sleep(100 * time.Millisecond)

	fmt.Println("Program complete")
}
