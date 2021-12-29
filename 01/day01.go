package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func LoadInput(filename string) []int {
	var input []int
	//Open file and create scanner
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		element, _ := strconv.Atoi(scanner.Text())

		input = append(input, element)
	}

	return input
}

func CalculateDepthIncreases(input []int) int {
	var counter int = 0
	var previous int

	for i, element := range input {
		if i == 0 {
			previous = element
		} else {
			if element > previous {
				counter++
			}
			previous = element
		}
	}

	return counter
}

func CalculateSlidingWindowDepthIncreases(input []int) int {
	var counter int = 0
	var previous [3]int

	for i, _ := range input {
		//Fill up First element
		if i < 3 {
			previous[i] = input[i]
		} else {
			//Main Calculation
			sum := input[i] + input[i-1] + input[i-2]
			previousSum := previous[0] + previous[1] + previous[2]
			if sum > previousSum {
				counter++
			}
			//Update previous
			previous[0] = previous[1]
			previous[1] = previous[2]
			previous[2] = input[i]
		}
	}

	return counter
}

func main() {
	input := LoadInput("input")

	fmt.Println(CalculateDepthIncreases(input))

	fmt.Println(CalculateSlidingWindowDepthIncreases(input))

}
