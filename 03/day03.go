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
	var counter [12]int

	//Open file and create scanner
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		s := strings.Split(scanner.Text(), "")

		for i := 0; i < 12; i++ {
			bit, _ := strconv.Atoi(s[i])

			if bit == 1 {
				counter[i]++
			} else {
				counter[i]--
			}
		}
	}

	//Evaluate counters
	for i := 0; i < 12; i++ {
		if counter[i] > 0 {
			counter[i] = 1
		} else {
			counter[i] = 0
		}
	}

	//Print final binary
	fmt.Println(counter)

	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()
}
