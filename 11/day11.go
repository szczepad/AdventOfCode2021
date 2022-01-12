package main

import (
	"bufio"
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

func IncreaseEnergy(input []int) []int {
	output := input

	for i, _ := range output {
		output[i]++
	}

	return output
}

func DoFlash(input []int, position int, cols int) []int {
	output := input

	x, y := ConvertPositionToCoordinates(position, cols)

	//Increase Left
	if (x - 1) >= 0 {
		output[ConvertCoordinatesToPosition(x-1, y, cols)]++
	}

	//Increase Right
	if (x + 1) < cols {
		output[ConvertCoordinatesToPosition(x+1, y, cols)]++
	}

	//Increase Top
	if (y - 1) >= 0 {
		output[ConvertCoordinatesToPosition(x, y-1, cols)]++
	}

	//Increase Bottom
	if ConvertCoordinatesToPosition(x, y+1, cols) < len(input) {
		output[ConvertCoordinatesToPosition(x, y+1, cols)]++
	}

	//Increase Top-Left
	if x-1 >= 0 && ConvertCoordinatesToPosition(x-1, y-1, cols) >= 0 {
		output[ConvertCoordinatesToPosition(x-1, y-1, cols)]++
	}

	//Increase Top-Right
	if (x+1) < cols && (y-1) >= 0 {
		output[ConvertCoordinatesToPosition(x+1, y-1, cols)]++
	}

	//Increase Bottom-Left
	if (x-1) >= 0 && ConvertCoordinatesToPosition(x-1, y+1, cols) < len(input) {
		output[ConvertCoordinatesToPosition((x-1), (y+1), cols)]++
	}

	//Increase Bottom-Right
	if (x+1) < cols && ConvertCoordinatesToPosition(x+1, y+1, cols) < len(input) {
		output[ConvertCoordinatesToPosition((x+1), (y+1), cols)]++
	}

	return output
}

func ConvertPositionToCoordinates(position int, cols int) (x int, y int) {
	x = position % cols
	y = (position - x) / cols

	return x, y
}

func ConvertCoordinatesToPosition(x int, y int, cols int) int {
	return (y * cols) + x
}

func HasAlreadyFlashed(flashedElements []int, toCheck int) bool {
	alreadyFlashed := false

	for _, elem := range flashedElements {
		if elem == toCheck {
			alreadyFlashed = true
		}
	}

	return alreadyFlashed
}

func ExecuteCycle(input []int, cols int) ([]int, int) {
	output := IncreaseEnergy(input)
	numberOfFlashes := 0

	allFlashed := false
	var flashed []int

	for !allFlashed {
		//Assume that it's not necessary to check one more time
		allFlashed = true

		for i, elem := range output {
			//Check if element is eligible for flash
			if elem > 9 && !HasAlreadyFlashed(flashed, i) {
				output = DoFlash(output, i, cols)
				flashed = append(flashed, i)
				numberOfFlashes++
				//Board has changed. Another Check necessary
				allFlashed = false
			}
		}

		for _, i := range flashed {
			output[i] = 0
		}
	}

	return output, numberOfFlashes
}

func CalculateFlashesAfterN(input []int, n int, cols int) int {
	totalFlashes := 0
	output := input

	for i := 0; i < n-1; i++ {
		tmpOutput, tmpFlashes := ExecuteCycle(output, cols)
		output = tmpOutput
		totalFlashes += tmpFlashes
	}

	return totalFlashes
}
