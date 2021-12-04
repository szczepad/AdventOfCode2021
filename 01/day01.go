package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var counter int = 0
	var previous int

	//Open file and create scanner
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	//Get initial value
	scanner.Scan()
	previous, _ = strconv.Atoi(scanner.Text())

	for scanner.Scan() {
		current, _ := strconv.Atoi(scanner.Text())

		if current > previous {
			counter++
		}
		previous = current
	}

	fmt.Println(counter)

	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()
}
