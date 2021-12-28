package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadInput(t *testing.T) {
	assert.Equal(t, []int{
		16,
		1,
		2,
		0,
		4,
		2,
		7,
		1,
		2,
		14,
	}, LoadInput("input_test"))
}

func TestGetMedian(t *testing.T) {
	assert.Equal(t, 2, GetMedian([]int{
		16,
		1,
		2,
		0,
		4,
		2,
		7,
		1,
		2,
		14,
	}))
}

func TestGetMean(t *testing.T) {
	lower, higher := GetMean([]int{
		16,
		1,
		2,
		0,
		4,
		2,
		7,
		1,
		2,
		14,
	})

	assert.Equal(t, 4, lower)
	assert.Equal(t, 5, higher)
}

func TestGetLinearTotalDistance(t *testing.T) {
	assert.Equal(t, 37, GetLinearTotalDistance([]int{
		16,
		1,
		2,
		0,
		4,
		2,
		7,
		1,
		2,
		14,
	}, 2))
}

func TestGetTotalDistance(t *testing.T) {
	assert.Equal(t, 168, GetTotalDistance([]int{
		16,
		1,
		2,
		0,
		4,
		2,
		7,
		1,
		2,
		14,
	}, 5))
}
