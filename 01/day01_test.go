package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadInput(t *testing.T) {
	assert.Equal(t, []int{
		199,
		200,
		208,
		210,
		200,
		207,
		240,
		269,
		260,
		263,
	}, LoadInput("input_test"))
}

func TestCalculateDepthIncreases(t *testing.T) {
	assert.Equal(t, 7, CalculateDepthIncreases([]int{
		199,
		200,
		208,
		210,
		200,
		207,
		240,
		269,
		260,
		263,
	}))
}

func TestCalculateSlidingWindowDepthIncreases(t *testing.T) {
	assert.Equal(t, 5, CalculateSlidingWindowDepthIncreases([]int{
		199,
		200,
		208,
		210,
		200,
		207,
		240,
		269,
		260,
		263,
	}))
}

func TestIsLowPoint(input []int, position int, cols int) {

}
