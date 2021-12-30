package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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

func IsCorrupted(exp []string) bool {
	if GetCorruptedSymbol(exp) == "" {
		return false
	}
	return true
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

func CalculateMissingScore(symbols []string) int {
	score := 0

	for _, elem := range symbols {
		score *= 5
		if elem == ")" {
			score += 1
		} else if elem == "]" {
			score += 2
		} else if elem == "}" {
			score += 3
		} else if elem == ">" {
			score += 4
		}
	}

	return score
}

func GetUncorruptedLines(input []string) []string {
	var validInputs []string

	for _, line := range input {
		if !IsCorrupted(LineToSlice(line)) {
			validInputs = append(validInputs, line)
		}
	}

	return validInputs
}

func GetMissingSymbols(exp []string) string {
	var stack []string

	var missing []string

	//Execute the Expression. Result after the loop should be open
	//Brackets with missing closing brackets
	for _, elem := range exp {
		if IsOpener(elem) {
			stack = append(stack, elem)
		} else {
			//Check if elem matches the currently needed opener
			if IsMatching(stack[len(stack)-1], elem) {
				stack = stack[:len(stack)-1] //Pop the Slice
			}
		}
	}

	revertedStack := RevertStack(stack)
	for _, elem := range revertedStack {
		missing = append(missing, GetClosingSymbol(elem))
	}

	return strings.Join(missing, "")
}

func RevertStack(stack []string) []string {
	var reverted []string

	for i := len(stack) - 1; i >= 0; i-- {
		reverted = append(reverted, stack[i])
	}

	return reverted
}

func GetClosingSymbol(symbol string) string {
	if symbol == "{" {
		return "}"
	} else if symbol == "(" {
		return ")"
	} else if symbol == "[" {
		return "]"
	} else if symbol == "<" {
		return ">"
	} else {
		return ""
	}
}

func main() {
	input := LoadFile("input")

	//Part 1
	var corruptedSymbols []string

	for _, line := range input {
		corruptedSymbols = append(corruptedSymbols, GetCorruptedSymbol(LineToSlice(line)))
	}

	fmt.Println(CalculateErrorScore(corruptedSymbols))

	//Part 2
	validInputs := GetUncorruptedLines(input)
	var missingSymbols []string
	var score []int

	//Get Missing Symbols
	for _, line := range validInputs {
		missingSymbols = append(missingSymbols, GetMissingSymbols(LineToSlice(line)))
	}

	//Calculate Scores
	for _, line := range missingSymbols {
		score = append(score, CalculateMissingScore(LineToSlice(line)))
	}

	//Sort Scores and output middle
	sort.Ints(score)
	fmt.Println(score[(len(score)-1)/2])
}
