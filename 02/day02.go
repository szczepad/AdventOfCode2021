package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var horizontal int = 0
	var depth int = 0

	//Open file and create scanner
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")

		distance, _ := strconv.Atoi(s[1])

		switch s[0] {
		case "forward":
			horizontal += distance
		case "up":
			depth -= distance
		case "down":
			depth += distance
		}
	}

	fmt.Println(horizontal * depth)

	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()
}
