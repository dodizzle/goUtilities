package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getEnvs(fileName string) map[string]string {
	vars := map[string]string{}
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines := strings.Split(scanner.Text(), "=")
		vars[lines[0]] = lines[1]
	}
	return vars
}

func printVars(vars map[string]string) {
	for k, v := range vars {
		fmt.Printf("%s=%s\n", k, v)
	}
}
func main() {
	vars := getEnvs("envs.txt")
	printVars(vars)
}
