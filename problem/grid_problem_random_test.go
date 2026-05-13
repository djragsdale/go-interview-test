package problem

import (
	"math/rand"
	"strings"
	"testing"
)

// Implemented this form the java tests

func TestSolveGridRandomPatternExists(t *testing.T) {
	solver := GridProblemImpl{}

	rng := rand.New(rand.NewSource(42))

	for i := 0; i < 100; i++ {
		grid := generateGrid(rng, 10, 10)
		pattern := generatePatternThatExists(t, rng, grid)

		got := solver.SolveGrid(grid, pattern)

		if !got {
			t.Fatalf(
				"SolveGrid() = false, want true\nGrid:\n%s\nPattern:\n%s",
				strings.Join(grid, "\n"),
				strings.Join(pattern, "\n"),
			)
		}
	}
}

func TestSolveGridRandomPatternDoesNotExist(t *testing.T) {
	solver := GridProblemImpl{}

	rng := rand.New(rand.NewSource(99))

	for i := 0; i < 100; i++ {
		grid := generateGrid(rng, 100, 100)

		// Use letters instead of digits so this pattern cannot appear in a
		// digit-only generated grid.
		pattern := generateLetterPattern(rng, 5, 5)

		got := solver.SolveGrid(grid, pattern)

		if got {
			t.Fatalf(
				"SolveGrid() = true, want false\nGrid:\n%s\nPattern:\n%s",
				strings.Join(grid, "\n"),
				strings.Join(pattern, "\n"),
			)
		}
	}
}

func TestSolveGridRandomLargePatternExists(t *testing.T) {
	solver := GridProblemImpl{}

	rng := rand.New(rand.NewSource(123))

	grid := generateGrid(rng, 500, 500)
	pattern := generatePatternThatExists(t, rng, grid)

	got := solver.SolveGrid(grid, pattern)

	if !got {
		t.Fatalf(
			"SolveGrid() = false, want true for large generated grid\nPattern:\n%s",
			strings.Join(pattern, "\n"),
		)
	}
}

func generateGrid(rng *rand.Rand, width, height int) []string {
	grid := make([]string, height)

	for row := 0; row < height; row++ {
		var builder strings.Builder
		builder.Grow(width)

		for col := 0; col < width; col++ {
			builder.WriteByte(byte('0' + rng.Intn(10)))
		}

		grid[row] = builder.String()
	}

	return grid
}

func generatePatternThatExists(t *testing.T, rng *rand.Rand, grid []string) []string {
	t.Helper()

	if len(grid) == 0 {
		t.Fatal("grid cannot be empty")
	}

	gridHeight := len(grid)
	gridWidth := len(grid[0])

	if gridWidth == 0 {
		t.Fatal("grid rows cannot be empty")
	}

	// Force dimensions to be at least 1.
	patternHeight := rng.Intn(gridHeight) + 1
	patternWidth := rng.Intn(gridWidth) + 1

	startRow := rng.Intn(gridHeight - patternHeight + 1)
	startCol := rng.Intn(gridWidth - patternWidth + 1)

	pattern := make([]string, patternHeight)

	for row := 0; row < patternHeight; row++ {
		sourceRow := grid[startRow+row]
		pattern[row] = sourceRow[startCol : startCol+patternWidth]
	}

	return pattern
}

func generateLetterPattern(rng *rand.Rand, width, height int) []string {
	pattern := make([]string, height)

	for row := 0; row < height; row++ {
		var builder strings.Builder
		builder.Grow(width)

		for col := 0; col < width; col++ {
			builder.WriteByte(byte('a' + rng.Intn(26)))
		}

		pattern[row] = builder.String()
	}

	return pattern
}
