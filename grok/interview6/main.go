package main

import "fmt"

func printFinal(sumOfRequests map[string]int) {
	fmt.Println("Request distribution:")
	for k, v := range sumOfRequests {
		fmt.Printf("server:%s || %d requests\n", k, v)
	}
}
func main() {
	//Write a Go function that simulates a simple load balancer.
	// It takes a list of backend server addresses and distributes incoming 'requests'
	// (represented as integers from 0 to N) across the servers using a round-robin strategy.
	//  Return a map showing how many requests each server handled.
	sumOfRequests := map[string]int{}
	backEndServers := []string{"10.10.1.1", "10.10.1.2", "10.10.1.3", "10.10.1.4", "10.10.1.5", "10.10.1.6"}
	// initialize reqeust count to be zero
	for l := 0; l < len(backEndServers); l++ {
		sumOfRequests[backEndServers[l]] = 0
	}

	for i := 0; i < 100; i++ {
		serverIndex := i % len(backEndServers)
		f := backEndServers[serverIndex]
		sumOfRequests[f] = sumOfRequests[f] + 1

	}
	printFinal(sumOfRequests)
}
