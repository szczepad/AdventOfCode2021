package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetOutputsFromFile(t *testing.T) {
	assert.Equal(t, []string{
		"fdgacbe cefdb cefbgd gcbe",
		"fcgedb cgb dgebacf gc",
		"cg cg fdcagb cbg",
		"efabcd cedba gadfec cb",
		"gecf egdcabf bgf bfgea",
		"gebdcfa ecba ca fadegcb",
		"cefg dcbef fcge gbcadfe",
		"ed bcgafe cdgba cbgef",
		"gbdfcae bgc cg cgb",
		"fgae cfgab fg bagce",
	}, GetOutputsFromFile("input_test"))
}

func TestIsOne(t *testing.T) {
	assert.Equal(t, true, IsOne("fg"))
}
func TestIsFour(t *testing.T) {
	assert.Equal(t, true, IsFour("gcbe"))
}
func TestIsSeven(t *testing.T) {
	assert.Equal(t, true, IsSeven("cgb"))
}

func TestIsEight(t *testing.T) {
	assert.Equal(t, true, IsEight("gbdfcae"))
}
