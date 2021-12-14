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
	//Open file and create scanner
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	//Get first line and construct an array with drawn numbers
	scanner.Scan()
	drawsString := strings.Split(scanner.Text(), ",")
	var draws []int
	for _, draw := range drawsString {
		tmp, _ := strconv.Atoi(draw)
		draws = append(draws, tmp)
	}

	fmt.Println(draws)
	//Skip empty line.
	scanner.Scan()

	//Construct the Bingo-Boards
	var bingoBoards [][]int
	var stringBuf [5]string
	rowIterator := 0

	for scanner.Scan() {
		if scanner.Text() == "" {
			currentBoardString := strings.Fields(strings.Join(stringBuf[0:5], " "))
			var currentBoard []int
			for _, cell := range currentBoardString {
				tmp, _ := strconv.Atoi(cell)
				currentBoard = append(currentBoard, tmp)
			}
			bingoBoards = append(bingoBoards, currentBoard)
			rowIterator = 0

		} else {
			stringBuf[rowIterator] = scanner.Text()
			rowIterator++
		}
	}

	for _, draw := range draws {
		for _, board := range bingoBoards {
			for i, cell := range board {
				if draw == cell {
					board[i] = 0
					if hasWon(board) {
						fmt.Println(calculateSum(board, draw))
						os.Exit(0)
					}
				}
			}
		}
	}

	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()
}

func hasWon(board []int) bool {
	var win bool = true

	for i := 0; i < 5; i++ {
		if board[i] != 0 {
			win = false
		}
	}
	if win {
		return win
	}

	win = true
	for i := 5; i < 10; i++ {
		if board[i] != 0 {
			win = false
		}
	}
	if win {
		return win
	}

	win = true
	for i := 10; i < 15; i++ {
		if board[i] != 0 {
			win = false
		}
	}
	if win {
		return win
	}

	win = true
	for i := 15; i < 20; i++ {
		if board[i] != 0 {
			win = false
		}
	}
	if win {
		return win
	}

	win = true
	for i := 20; i < 25; i++ {
		if board[i] != 0 {
			win = false
		}
	}
	if win {
		return win
	}

	win = true
	for i := 0; i < 25; i += 5 {
		if board[i] != 0 {
			win = false
		}
	}
	if win {
		return win
	}

	win = true
	for i := 1; i < 25; i += 5 {
		if board[i] != 0 {
			win = false
		}
	}
	if win {
		return win
	}

	win = true
	for i := 2; i < 25; i += 5 {
		if board[i] != 0 {
			win = false
		}
	}
	if win {
		return win
	}

	win = true
	for i := 3; i < 25; i += 5 {
		if board[i] != 0 {
			win = false
		}
	}
	if win {
		return win
	}

	win = true
	for i := 4; i < 25; i += 5 {
		if board[i] != 0 {
			win = false
		}
	}
	if win {
		return win
	}

	return false

}

func calculateSum(board []int, draw int) int {
	var sum int = 0
	for _, x := range board {
		sum += x
	}

	return sum * draw
}

func printBoard(board []int) {
	for i, cell := range board {
		if i%5 == 0 {
			fmt.Printf("\n")
		}
		fmt.Printf("%d ", cell)
	}
	fmt.Printf("\n")
}
