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
	if (y+1)+x < len(input) {
		output[ConvertCoordinatesToPosition(x, y+1, cols)]++
	}

	//Increase Top-Left
	if ConvertCoordinatesToPosition(x-1, y-1, cols) >= 0 {
		output[ConvertCoordinatesToPosition(x-1, y-1, cols)]++
	}

	//Increase Top-Right
	if (x+1) < cols && (y-1) >= 0 {
		output[ConvertCoordinatesToPosition(x+1, y-1, cols)]++
	}

	//Increase Bottom-Left
	if (x-1) >= 0 && (y+1)+x < len(input) {
		output[ConvertCoordinatesToPosition((x-1), (y+1), cols)]++
	}

	//Increase Bottom-Right
	if (x+1) < cols && (y+1)+x < len(input) {
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
