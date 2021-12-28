package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func LoadInput(filePath string) []int {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	stringSlice := strings.Split(scanner.Text(), ",")
	var intSlice []int

	for _, element := range stringSlice {
		tmp, _ := strconv.Atoi(element)
		intSlice = append(intSlice, tmp)
	}

	return intSlice
}

func GetMedian(items []int) int {
	sort.Ints(items)

	mNumber := len(items) / 2

	if len(items)%2 == 1 {
		return items[mNumber]
	}

	return (items[mNumber-1] + items[mNumber]) / 2
}

func GetMean(items []int) (lower int, higher int) {
	total := 0

	for _, element := range items {
		total += element
	}

	mean := float64(total) / float64(len(items))

	return int(math.Floor(mean)), int(math.Ceil(mean))
}

func GetLinearTotalDistance(items []int, goal int) int {
	var totalDistance float64

	for _, element := range items {
		totalDistance += math.Abs(float64(element) - float64(goal))
	}
	return int(totalDistance)
}

func GetTotalDistance(items []int, goal int) int {
	var total float64

	for _, element := range items {
		n := math.Abs(float64(goal) - float64(element))
		total += n * (n + 1) / 2
	}

	return int(total)
}

func main() {
	input := LoadInput("input")
	goal := GetMedian(input)

	fmt.Println(GetLinearTotalDistance(input, goal))

	// Part 2
	lower, higher := GetMean(input)
	result1 := GetTotalDistance(input, lower)
	result2 := GetTotalDistance(input, higher)

	if result1 < result2 {
		fmt.Println(result1)
	} else {
		fmt.Println(result2)
	}
}
