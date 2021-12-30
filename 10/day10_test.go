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

func TestIsCorrupted(t *testing.T) {
	input1 := []string{
		"{", "(", "[", "(", "<", "{", "}", "[", "<", ">", "[", "]", "}",
		">", "{", "[", "]", "{", "[", "(", "<", "(", ")", ">",
	}
	input2 := []string{
		"[", "(", "{", "(", "<", "(", "(", ")", ")", "[", "]", ">", "[",
		"[", "{", "[", "]", "{", "<", "(", ")", "<", ">", ">",
	}

	assert.Equal(t, true, IsCorrupted(input1))
	assert.Equal(t, false, IsCorrupted(input2))
}

func TestGetUncorruptedLines(t *testing.T) {
	input := []string{
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
	}

	output := []string{
		"[({(<(())[]>[[{[]{<()<>>",
		"[(()[<>])]({[<{<<[]>>(",
		"(((({<>}<{<{<>}{[]{[]{}",
		"{<[[]]>}<{[{[{[]{()[[[]",
		"<{([{{}}[<[[[<>{}]]]>[]]",
	}

	assert.Equal(t, output, GetUncorruptedLines(input))
}

func TestGetMissingSymbols(t *testing.T) {
	input := []string{
		"[", "(", "{", "(", "<", "(", "(", ")", ")", "[", "]",
		">", "[", "[", "{", "[", "]", "{", "<", "(", ")", "<", ">", ">",
	}
	output := "}}]])})]"

	assert.Equal(t, output, GetMissingSymbols(input))
}

func TestRevertStack(t *testing.T) {
	input := []string{"[", "(", "{", "(", "[", "[", "{", "{"}
	output := []string{"{", "{", "[", "[", "(", "{", "(", "["}

	assert.Equal(t, output, RevertStack(input))
}

func TestGetClosingSymbol(t *testing.T) {
	assert.Equal(t, "]", GetClosingSymbol("["))
}

func TestCalculateMissingScore(t *testing.T) {
	assert.Equal(t, 294, CalculateMissingScore([]string{"]", ")", "}", ">"}))
	assert.Equal(t, 288957, CalculateMissingScore([]string{"}", "}", "]", "]", ")", "}", ")", "]"}))
}
