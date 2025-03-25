package main

import "fmt"

type Server struct {
	Name   string
	IP     string
	Status string
}

func (s Server) isOnline() bool {
	if s.Status == "online" {
		return true
	}
	return false
}

func (s Server) Print() {
	fmt.Printf("Server: %s IP: %s Status: %s\n", s.Name, s.IP, s.Status)
}

func main() {
	//Write a Go program that models a simple server inventory system.
	// Create a Server struct to represent a server with fields for its name, IP address,
	// and status ('online' or 'offline'). Add a method to the Server struct that checks if the server is online.
	// Then, write a function that takes a slice of Server structs and returns a count of how many are online.
	inventory := []Server{
		{Name: "web1", IP: "192.168.1.10", Status: "online"},
		{Name: "db1", IP: "192.168.1.20", Status: "offline"},
		{Name: "app1", IP: "192.168.1.30", Status: "online"},
		{Name: "cache1", IP: "192.168.1.40", Status: "offline"},
	}
	online := 0
	for _, v := range inventory {
		if v.isOnline() {
			online++
		}
	}
	fmt.Printf("%d servers are online\n", online)
	for _, v := range inventory {
		v.Print()
	}
}
