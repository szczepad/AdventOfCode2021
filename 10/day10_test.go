package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadFile(t *testing.T) {
	assert.Equal(t, []string{
		"[({(<(())[]>[[{[]{<()<>>",
		"[(()[<>])]({[<{<<[]>>(",
		"{([(<{}[<>[]}>{[]{[(<()>",
		"(((({<>}<{<{<>}{[]{[]{}",
		"[[<[([]))<([[{}[[()]]]",
		"[{[{({}]{}}([{[{{{}}([]",
		"{<[[]]>}<{[{[{[]{()[[[]",
		"[<(<(<(<{}))><([]([]()",
		"<{([([[(<>()){}]>(<<{{",
		"<{([{{}}[<[[[<>{}]]]>[]]",
	}, LoadFile("input_test"))
}

func TestLineToSlice(t *testing.T) {
	assert.Equal(t, []string{
		"[",
		"(",
		"{",
		"(",
		"<",
		"(",
		"(",
		")",
		")",
		"[",
		"]",
		">",
		"[",
		"[",
		"{",
		"[",
		"]",
		"{",
		"<",
		"(",
		")",
		"<",
		">",
		">",
	}, LineToSlice("[({(<(())[]>[[{[]{<()<>>"))
}

func TestIsOpener(t *testing.T) {
	assert.Equal(t, true, IsOpener("["))
	assert.Equal(t, false, IsOpener(">"))
}

func TestIsCloser(t *testing.T) {
	assert.Equal(t, false, IsCloser("["))
	assert.Equal(t, true, IsCloser(">"))
}

func TestIsMatching(t *testing.T) {
	assert.Equal(t, true, IsMatching("[", "]"))
	assert.Equal(t, false, IsMatching("(", ">"))
}

func TestGetCorruptedSymbol(t *testing.T) {
	input := []string{
		"{",
		"(",
		"[",
		"(",
		"<",
		"{",
		"}",
		"[",
		"<",
		">",
		"[",
		"]",
		"}",
		">",
		"{",
		"[",
		"]",
		"{",
		"[",
		"(",
		"<",
		"(",
		")",
		">",
	}
	assert.Equal(t, "}", GetCorruptedSymbol(input))
}

func TestCalculateErrorScore(t *testing.T) {
	input := []string{
		"}",
		")",
		"]",
		")",
		">",
	}

	assert.Equal(t, 26397, CalculateErrorScore(input))
}
