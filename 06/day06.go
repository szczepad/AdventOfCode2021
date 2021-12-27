package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func LoadInput(filePath string) []int {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	stringSlice := strings.Split(scanner.Text(), ",")
	var intSlice []int

	for _, element := range stringSlice {
		tmp, _ := strconv.Atoi(element)
		intSlice = append(intSlice, tmp)
	}

	return intSlice
}

func DoCycles(input []int, iterations int) []int {
	if iterations == 0 {
		return input
	} else {
		for i, cell := range input {
			if cell == 0 {
				input[i] = 6
				input = append(input, 8)
			} else {
				input[i]--
			}

		}
		return DoCycles(input, iterations-1)
	}
}

func CountLanternfish(input []int) int {
	return len(input)
}

func main() {
	input := LoadInput("input")

	fmt.Println(CountLanternfish(DoCycles(input, 80)))
}
