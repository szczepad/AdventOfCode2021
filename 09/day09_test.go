package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadFile(t *testing.T) {
	input, cols := LoadFile("input_test")
	assert.Equal(t, []int{
		2, 1, 9, 9, 9, 4, 3, 2, 1, 0,
		3, 9, 8, 7, 8, 9, 4, 9, 2, 1,
		9, 8, 5, 6, 7, 8, 9, 8, 9, 2,
		8, 7, 6, 7, 8, 9, 6, 7, 8, 9,
		9, 8, 9, 9, 9, 6, 5, 6, 7, 8,
	}, input)

	assert.Equal(t, 10, cols)
}

func TestIsLowPoint(t *testing.T) {
	input := []int{
		2, 1, 9, 9, 9, 4, 3, 2, 1, 0,
		3, 9, 8, 7, 8, 9, 4, 9, 2, 1,
		9, 8, 5, 6, 7, 8, 9, 8, 9, 2,
		8, 7, 6, 7, 8, 9, 6, 7, 8, 9,
		9, 8, 9, 9, 9, 6, 5, 6, 7, 8,
	}
	cols := 10

	assert.Equal(t, true, IsLowPoint(input, cols, 1))
	assert.Equal(t, false, IsLowPoint(input, cols, 0))
	assert.Equal(t, false, IsLowPoint(input, cols, 2))
	assert.Equal(t, true, IsLowPoint(input, cols, 9))
	assert.Equal(t, false, IsLowPoint(input, cols, 45))
}

func TestCalculateRiskLevel(t *testing.T) {
	input := []int{
		2, 1, 9, 9, 9, 4, 3, 2, 1, 0,
		3, 9, 8, 7, 8, 9, 4, 9, 2, 1,
		9, 8, 5, 6, 7, 8, 9, 8, 9, 2,
		8, 7, 6, 7, 8, 9, 6, 7, 8, 9,
		9, 8, 9, 9, 9, 6, 5, 6, 7, 8,
	}
	cols := 10

	assert.Equal(t, 15, CalculateRiskLevel(input, cols))
}
