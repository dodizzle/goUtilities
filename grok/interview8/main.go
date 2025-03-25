package main

import (
	"fmt"
	"strconv"
	"strings"
)

func checkValid(line string) bool {
	pieces := strings.Split(line, ".")
	if len(pieces) != 4 {
		return false
	}
	for _, v := range pieces {
		num, err := strconv.Atoi(v)
		if err != nil || num < 0 || num > 255 {
			return false
		}
	}
	return true
}

func main() {
	//Write a Go function that takes a list of IP addresses (as strings) and checks if each one is valid IPv4 format (e.g., '192.168.1.1').
	// Return a slice of only the valid IPs.
	// You don’t need to verify if they’re routable, just that they follow the correct format:
	// four numbers (0-255) separated by dots.
	addresses := []string{"1.1.1.111", "2234.45.7.3", "192.168.51.1", "1.1.1"}
	for _, v := range addresses {
		if checkValid(v) {
			fmt.Println(v)
		}
	}
}
