package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadInput(t *testing.T) {
	assert.Equal(t, []string{
		"0,9 -> 5,9",
		"8,0 -> 0,8",
		"9,4 -> 3,4",
		"2,2 -> 2,1",
		"7,0 -> 7,4",
		"6,4 -> 2,0",
		"0,9 -> 2,9",
		"3,4 -> 1,4",
		"0,0 -> 8,8",
		"5,5 -> 8,2",
	}, LoadInput("input_test"))
}

func TestConvertLineToCoordinates(t *testing.T) {
	assert.Equal(t, []int{
		0,
		9,
		5,
		9,
	}, ConvertLineToCoordinates("0,9 -> 5,9"))
}

func TestIsDiagonal(t *testing.T) {
	assert.Equal(t, false, IsDiagonal([]int{0, 9, 5, 9}))
}

func TestDrawLine(t *testing.T) {
	const rows = 10
	const cols = 10

	var board = make([]int, rows*cols)

	input := []int{0, 9, 5, 9}

	DrawLine(board, rows, input)

	//Check whether the line was drawn
	assert.Equal(t, 1, board[0*rows+9])
	assert.Equal(t, 1, board[1*rows+9])
	assert.Equal(t, 1, board[2*rows+9])
	assert.Equal(t, 1, board[3*rows+9])
	assert.Equal(t, 1, board[4*rows+9])
	assert.Equal(t, 1, board[5*rows+9])
}
func TestGetOverlaps(t *testing.T) {

}
