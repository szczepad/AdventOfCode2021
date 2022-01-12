package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadFile(t *testing.T) {
	input, cols := LoadFile("input_test")
	assert.Equal(t, []int{
		5, 4, 8, 3, 1, 4, 3, 2, 2, 3,
		2, 7, 4, 5, 8, 5, 4, 7, 1, 1,
		5, 2, 6, 4, 5, 5, 6, 1, 7, 3,
		6, 1, 4, 1, 3, 3, 6, 1, 4, 6,
		6, 3, 5, 7, 3, 8, 5, 4, 7, 8,
		4, 1, 6, 7, 5, 2, 4, 6, 4, 5,
		2, 1, 7, 6, 8, 4, 1, 7, 2, 1,
		6, 8, 8, 2, 8, 8, 1, 1, 3, 4,
		4, 8, 4, 6, 8, 4, 8, 5, 5, 4,
		5, 2, 8, 3, 7, 5, 1, 5, 2, 6,
	}, input)

	assert.Equal(t, 10, cols)
}

func TestIncreaseEnergy(t *testing.T) {
	input := []int{
		5, 4, 8, 3, 1, 4, 3, 2, 2, 3,
		2, 7, 4, 5, 8, 5, 4, 7, 1, 1,
		5, 2, 6, 4, 5, 5, 6, 1, 7, 3,
		6, 1, 4, 1, 3, 3, 6, 1, 4, 6,
		6, 3, 5, 7, 3, 8, 5, 4, 7, 8,
		4, 1, 6, 7, 5, 2, 4, 6, 4, 5,
		2, 1, 7, 6, 8, 4, 1, 7, 2, 1,
		6, 8, 8, 2, 8, 8, 1, 1, 3, 4,
		4, 8, 4, 6, 8, 4, 8, 5, 5, 4,
		5, 2, 8, 3, 7, 5, 1, 5, 2, 6,
	}
	output := []int{
		6, 5, 9, 4, 2, 5, 4, 3, 3, 4,
		3, 8, 5, 6, 9, 6, 5, 8, 2, 2,
		6, 3, 7, 5, 6, 6, 7, 2, 8, 4,
		7, 2, 5, 2, 4, 4, 7, 2, 5, 7,
		7, 4, 6, 8, 4, 9, 6, 5, 8, 9,
		5, 2, 7, 8, 6, 3, 5, 7, 5, 6,
		3, 2, 8, 7, 9, 5, 2, 8, 3, 2,
		7, 9, 9, 3, 9, 9, 2, 2, 4, 5,
		5, 9, 5, 7, 9, 5, 9, 6, 6, 5,
		6, 3, 9, 4, 8, 6, 2, 6, 3, 7,
	}

	assert.Equal(t, output, IncreaseEnergy(input))
}

func TestConvertPositionToCoordinates(t *testing.T) {
	x, y := ConvertPositionToCoordinates(11, 10)

	assert.Equal(t, x, 1)
	assert.Equal(t, y, 1)
}

func TestConvertCoordinatesToPosition(t *testing.T) {
	assert.Equal(t, 11, ConvertCoordinatesToPosition(1, 1, 10))
}

func TestDoFlash(t *testing.T) {
	input1 := []int{
		6, 5, 9, 4, 2, 5, 4, 3, 3, 4,
		3, 8, 5, 6, 9, 6, 5, 8, 2, 2,
		6, 3, 7, 5, 6, 6, 7, 2, 8, 4,
		7, 2, 5, 2, 4, 4, 7, 2, 5, 7,
		7, 4, 6, 8, 4, 9, 6, 5, 8, 9,
		5, 2, 7, 8, 6, 3, 5, 7, 5, 6,
		3, 2, 8, 7, 9, 5, 2, 8, 3, 2,
		7, 9, 9, 3, 9, 9, 2, 2, 4, 5,
		5, 9, 5, 7, 9, 5, 9, 6, 6, 5,
		6, 3, 9, 4, 8, 6, 2, 6, 3, 7,
	}

	output1 := []int{
		6, 6, 9, 5, 2, 5, 4, 3, 3, 4,
		3, 9, 6, 7, 9, 6, 5, 8, 2, 2,
		6, 3, 7, 5, 6, 6, 7, 2, 8, 4,
		7, 2, 5, 2, 4, 4, 7, 2, 5, 7,
		7, 4, 6, 8, 4, 9, 6, 5, 8, 9,
		5, 2, 7, 8, 6, 3, 5, 7, 5, 6,
		3, 2, 8, 7, 9, 5, 2, 8, 3, 2,
		7, 9, 9, 3, 9, 9, 2, 2, 4, 5,
		5, 9, 5, 7, 9, 5, 9, 6, 6, 5,
		6, 3, 9, 4, 8, 6, 2, 6, 3, 7,
	}

	output2 := []int{
		6, 6, 9, 5, 2, 5, 4, 3, 3, 4,
		3, 9, 6, 7, 9, 6, 5, 8, 2, 2,
		6, 3, 7, 5, 6, 6, 7, 2, 8, 4,
		7, 2, 5, 2, 4, 4, 7, 2, 5, 7,
		7, 4, 6, 8, 4, 9, 6, 5, 8, 9,
		5, 2, 7, 8, 6, 3, 5, 7, 5, 6,
		3, 2, 8, 7, 9, 5, 2, 8, 3, 2,
		7, 9, 9, 3, 9, 9, 2, 2, 4, 5,
		5, 9, 5, 7, 9, 5, 9, 6, 7, 6,
		6, 3, 9, 4, 8, 6, 2, 6, 4, 7,
	}

	assert.Equal(t, output1, DoFlash(input1, 2, 10))
	assert.Equal(t, output2, DoFlash(input1, 99, 10))
}

func TesHasAlreadyFlashed(t *testing.T) {
	alreadyFlashed := []int{1, 2, 3, 4}

	assert.Equal(t, true, HasAlreadyFlashed(alreadyFlashed, 1))
	assert.Equal(t, false, HasAlreadyFlashed(alreadyFlashed, 0))
}
func TestExecuteCycle(t *testing.T) {
	input := []int{
		6, 5, 9, 4, 2, 5, 4, 3, 3, 4,
		3, 8, 5, 6, 9, 6, 5, 8, 2, 2,
		6, 3, 7, 5, 6, 6, 7, 2, 8, 4,
		7, 2, 5, 2, 4, 4, 7, 2, 5, 7,
		7, 4, 6, 8, 4, 9, 6, 5, 8, 9,
		5, 2, 7, 8, 6, 3, 5, 7, 5, 6,
		3, 2, 8, 7, 9, 5, 2, 8, 3, 2,
		7, 9, 9, 3, 9, 9, 2, 2, 4, 5,
		5, 9, 5, 7, 9, 5, 9, 6, 6, 5,
		6, 3, 9, 4, 8, 6, 2, 6, 3, 7,
	}

	expectedOutput := []int{
		8, 8, 0, 7, 4, 7, 6, 5, 5, 5,
		5, 0, 8, 9, 0, 8, 7, 0, 5, 4,
		8, 5, 9, 7, 8, 8, 9, 6, 0, 8,
		8, 4, 8, 5, 7, 6, 9, 6, 0, 0,
		8, 7, 0, 0, 9, 0, 8, 8, 0, 0,
		6, 6, 0, 0, 0, 8, 8, 9, 8, 9,
		6, 8, 0, 0, 0, 0, 5, 9, 4, 3,
		0, 0, 0, 0, 0, 0, 7, 4, 5, 6,
		9, 0, 0, 0, 0, 0, 0, 8, 7, 6,
		8, 7, 0, 0, 0, 0, 6, 8, 4, 8,
	}

	actualOutput, flashes := ExecuteCycle(input, 10)
	assert.Equal(t, expectedOutput, actualOutput)
	assert.Equal(t, 35, flashes)
}

func TestCalculateFlashesAfterN(t *testing.T) {
	input := []int{
		5, 4, 8, 3, 1, 4, 3, 2, 2, 3,
		2, 7, 4, 5, 8, 5, 4, 7, 1, 1,
		5, 2, 6, 4, 5, 5, 6, 1, 7, 3,
		6, 1, 4, 1, 3, 3, 6, 1, 4, 6,
		6, 3, 5, 7, 3, 8, 5, 4, 7, 8,
		4, 1, 6, 7, 5, 2, 4, 6, 4, 5,
		2, 1, 7, 6, 8, 4, 1, 7, 2, 1,
		6, 8, 8, 2, 8, 8, 1, 1, 3, 4,
		4, 8, 4, 6, 8, 4, 8, 5, 5, 4,
		5, 2, 8, 3, 7, 5, 1, 5, 2, 6,
	}

	actualFlashes := CalculateFlashesAfterN(input, 100, 10)

	assert.Equal(t, 1656, actualFlashes)
}
