package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func LoadFile(filePath string) []string {
	var input []string

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
	}

	return input
}

func LineToSlice(line string) []string {

	return strings.Split(line, "")
}

func IsOpener(item string) bool {
	if item == "{" || item == "(" || item == "[" || item == "<" {
		return true
	}
	return false
}

func IsCloser(item string) bool {
	if item == "}" || item == ")" || item == "]" || item == ">" {
		return true
	}
	return false
}

func IsMatching(opener string, item string) bool {
	if (opener == "{" && item == "}") || (opener == "(" && item == ")") ||
		(opener == "[" && item == "]") || (opener == "<" && item == ">") {
		return true
	}
	return false
}

func GetCorruptedSymbol(exp []string) string {
	var stack []string

	for _, elem := range exp {
		if IsOpener(elem) {
			stack = append(stack, elem)
		} else if IsCloser(elem) {
			//Check if elem matches the currently needed opener
			if IsMatching(stack[len(stack)-1], elem) {
				stack = stack[:len(stack)-1] //Pop the Slice
			} else {
				return elem
			}
		}
	}

	return ""
}

func CalculateErrorScore(errors []string) int {
	score := 0

	for _, elem := range errors {
		if elem == ")" {
			score += 3
		} else if elem == "]" {
			score += 57
		} else if elem == "}" {
			score += 1197
		} else if elem == ">" {
			score += 25137
		}
	}

	return score
}

func main() {
	input := LoadFile("input")

	var corruptedSymbols []string

	for _, line := range input {
		corruptedSymbols = append(corruptedSymbols, GetCorruptedSymbol(LineToSlice(line)))
	}

	fmt.Println(CalculateErrorScore(corruptedSymbols))
}
