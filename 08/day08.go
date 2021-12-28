package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func GetOutputsFromFile(filePath string) []string {
	var outputs []string

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		outputElement := strings.TrimSpace(strings.Split(line, "|")[1])
		outputs = append(outputs, outputElement)
	}

	return outputs
}

func IsOne(s string) bool {
	if len(s) == 2 {
		return true
	} else {
		return false
	}
}

func IsFour(s string) bool {
	if len(s) == 4 {
		return true
	} else {
		return false
	}
}

func IsSeven(s string) bool {
	if len(s) == 3 {
		return true
	} else {
		return false
	}
}

func IsEight(s string) bool {
	if len(s) == 7 {
		return true
	} else {
		return false
	}
}
