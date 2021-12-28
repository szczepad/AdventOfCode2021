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

func IsFour(s string) bool {
	if len(s) == 4 {
		return true
	} else {
		return false
	}
}
