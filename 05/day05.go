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

func main() {
	input := LoadInput("input")

	fmt.Println(input)
}
