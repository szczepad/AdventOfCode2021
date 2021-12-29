package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func LoadFile(filePath string) (input []int, cols int) {
	var inputTemp []int

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()
		cols = len(text)
		textChars := strings.Split(text, "")
		for _, element := range textChars {
			number, _ := strconv.Atoi(element)
			inputTemp = append(inputTemp, number)
		}
	}

	return inputTemp, cols
}

func IsLowPoint(input []int, cols int, position int) bool {
	//Convert Position in array to Coordinate
	x := position % cols
	y := (position - x) / cols

	//Get Current Position
	current := input[position]

	//Check tile to the left
	if (x - 1) >= 0 {
		if current >= input[(y*cols)+(x-1)] {
			return false
		}
	}
	//Check tile to the right
	if (x + 1) < cols {
		if current >= input[(y*cols)+(x+1)] {
			return false
		}
	}
	//Check top
	if (y - 1) >= 0 {
		if current >= input[((y-1)*cols)+x] {
			return false
		}
	}
	//Check bottom
	if ((y+1)*cols)+x < len(input) {
		if current >= input[((y+1)*cols)+x] {
			return false
		}
	}
	return true
}

func CalculateRiskLevel(input []int, cols int) int {
	riskLevel := 0

	for i, _ := range input {
		if IsLowPoint(input, cols, i) {
			riskLevel += 1 + input[i]
		}
	}
	return riskLevel
}

func main() {
	input, cols := LoadFile("input")

	fmt.Println(CalculateRiskLevel(input, cols))
}
