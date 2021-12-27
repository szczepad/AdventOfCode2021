package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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
	replacedString := strings.Replace(inputLine, " -> ", ",", 1)
	stringCoordinates := strings.Split(replacedString, ",")

	var intCoordinates []int

	for _, element := range stringCoordinates {
		tmp, _ := strconv.Atoi(element)
		intCoordinates = append(intCoordinates, tmp)
	}

	return intCoordinates
}

func IsDiagonal(coords []int) bool {
	if coords[0] == coords[2] || coords[1] == coords[3] {
		return false
	}
	return true
}

func DrawLine(board []int, rows int, input []int) {
	x1 := input[0]
	y1 := input[1]
	x2 := input[2]
	y2 := input[3]

	board[rows*x1+y1]++

	if x1 == x2 && y1 == y2 {
		return
	} else {
		if x1 < x2 {
			x1++
		} else if x1 > x2 {
			x1--
		}
		if y1 < y2 {
			y1++
		} else if y1 > y2 {
			y1--
		}

		DrawLine(board, rows, []int{
			x1,
			y1,
			x2,
			y2,
		})
	}
}

func CountOverlaps(board []int) int {
	var counter int = 0

	for _, cell := range board {
		if cell > 1 {
			counter++
		}
	}

	return counter
}

func main() {
	input := LoadInput("input")

	const rows = 999
	const cols = 999

	var board = make([]int, rows*cols)

	for _, line := range input {
		coords := ConvertLineToCoordinates(line)
		//Uncomment if-clause to get solution for Part A)
		//if !IsDiagonal(coords) {
		DrawLine(board, rows, coords)
		//}
	}

	fmt.Println(CountOverlaps(board))
}
