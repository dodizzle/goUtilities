package main

import (
	"math/rand"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func randomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func main() {
	type Person struct {
		name string
		age  int
	}

	personx := Person{}
	for range 10_000_000 {
		personx.name = randomString(10)
		personx.age = rand.Intn(100)
	}

}
