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

func ConvertLineToCoordinates(inputLine string) []int {
	return []int{1, 0, 0}
}

func main() {
	input := LoadInput("input")

	fmt.Println(input)
}
