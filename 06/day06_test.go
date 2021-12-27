package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadInput(t *testing.T) {
	assert.Equal(t, []int{
		3,
		4,
		3,
		1,
		2,
	}, LoadInput("input_test"))
}

func TestDoZeroCycles(t *testing.T) {
	assert.Equal(t, []int{
		3,
		4,
		3,
		1,
		2,
	}, DoCycles([]int{3, 4, 3, 1, 2}, 0))
}

func TestDoCycle(t *testing.T) {
	assert.Equal(t, []int{
		2,
		3,
		2,
		0,
		1,
	}, DoCycles([]int{3, 4, 3, 1, 2}, 1))
}

func TestDoCycleWithBirth(t *testing.T) {
	assert.Equal(t, []int{
		1,
		2,
		1,
		6,
		0,
		8,
	}, DoCycles([]int{3, 4, 3, 1, 2}, 2))
}

func TestCountLanternfishAfter18(t *testing.T) {
	finished := DoCycles([]int{3, 4, 3, 1, 2}, 18)

	assert.Equal(t, 26, CountLanternfish(finished))
}
