package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func LoadInput(filePath string) []string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	var input []string

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	return input
}

func main() {
	input := LoadInput("input")

	fmt.Println(input)
}
