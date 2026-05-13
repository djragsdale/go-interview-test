package problem

import (
	"strings"
	"testing"
)

func TestSolveGrid(t *testing.T) {
	solver := GridProblemImpl{}

	tests := []struct {
		name    string
		grid    []string
		pattern []string
		want    bool
	}{
		{
			name: "pattern exists in sample grid",
			grid: []string{
				"12930482091238127",
				"12037128741023801",
				"19208301741029388",
				"19238012938172401",
				"19123812038127415",
			},
			pattern: []string{
				"1920",
				"1923",
				"1912",
			},
			want: true,
		},
		{
			name: "pattern does not exist",
			grid: []string{
				"1234",
				"5678",
				"9012",
			},
			pattern: []string{
				"12",
				"90",
			},
			want: false,
		},
		{
			name: "pattern exists at top left",
			grid: []string{
				"abcd",
				"efgh",
				"ijkl",
			},
			pattern: []string{
				"ab",
				"ef",
			},
			want: true,
		},
		{
			name: "pattern exists at bottom right",
			grid: []string{
				"xxxxxx",
				"xxxxxx",
				"xxxxab",
				"xxxxcd",
			},
			pattern: []string{
				"ab",
				"cd",
			},
			want: true,
		},
		{
			name: "single cell pattern exists",
			grid: []string{
				"abc",
				"def",
				"ghi",
			},
			pattern: []string{
				"e",
			},
			want: true,
		},
		{
			name: "single cell pattern does not exist",
			grid: []string{
				"abc",
				"def",
				"ghi",
			},
			pattern: []string{
				"z",
			},
			want: false,
		},
		{
			name: "single row pattern exists",
			grid: []string{
				"abcdef",
				"ghijkl",
			},
			pattern: []string{
				"cde",
			},
			want: true,
		},
		{
			name: "single column pattern exists",
			grid: []string{
				"a123",
				"b456",
				"c789",
			},
			pattern: []string{
				"a",
				"b",
				"c",
			},
			want: true,
		},
		{
			name: "pattern equals entire grid",
			grid: []string{
				"abc",
				"def",
				"ghi",
			},
			pattern: []string{
				"abc",
				"def",
				"ghi",
			},
			want: true,
		},
		{
			name: "pattern taller than grid",
			grid: []string{
				"abc",
			},
			pattern: []string{
				"a",
				"b",
			},
			want: false,
		},
		{
			name: "pattern wider than grid row",
			grid: []string{
				"abc",
				"def",
			},
			pattern: []string{
				"abcd",
			},
			want: false,
		},
		{
			name:    "empty grid with non-empty pattern",
			grid:    []string{},
			pattern: []string{"a"},
			want:    false,
		},
		{
			name:    "empty pattern",
			grid:    []string{"abc"},
			pattern: []string{},
			want:    true,
		},
		{
			name: "ragged pattern is invalid",
			grid: []string{
				"abcdef",
				"ghijkl",
			},
			pattern: []string{
				"abc",
				"gh",
			},
			want: false,
		},
		{
			name: "repeated first pattern row but full pattern does not match",
			grid: []string{
				"aaaaaa",
				"bbbbbb",
				"aaaaaa",
				"cccccc",
			},
			pattern: []string{
				"aaa",
				"ddd",
			},
			want: false,
		},
		{
			name: "ragged grid still handles valid matching area",
			grid: []string{
				"abcde",
				"bc",
				"abcde",
			},
			pattern: []string{
				"abc",
				"abc",
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := solver.SolveGrid(tt.grid, tt.pattern)

			if got != tt.want {
				t.Fatalf("SolveGrid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolveGridLargeGridPatternExists(t *testing.T) {
	solver := GridProblemImpl{}

	const rows = 300
	const cols = 300

	grid := make([]string, rows)
	for i := range grid {
		grid[i] = strings.Repeat("0", cols)
	}

	pattern := []string{
		"12345",
		"67890",
		"abcde",
	}

	startRow := 250
	startCol := 240

	grid[startRow] = replaceAt(grid[startRow], startCol, pattern[0])
	grid[startRow+1] = replaceAt(grid[startRow+1], startCol, pattern[1])
	grid[startRow+2] = replaceAt(grid[startRow+2], startCol, pattern[2])

	got := solver.SolveGrid(grid, pattern)

	if !got {
		t.Fatalf("SolveGrid() = false, want true for large grid")
	}
}

func TestSolveGridLargeGridPatternDoesNotExist(t *testing.T) {
	solver := GridProblemImpl{}

	const rows = 300
	const cols = 300

	grid := make([]string, rows)
	for i := range grid {
		grid[i] = strings.Repeat("0", cols)
	}

	pattern := []string{
		"12345",
		"67890",
		"abcde",
	}

	got := solver.SolveGrid(grid, pattern)

	if got {
		t.Fatalf("SolveGrid() = true, want false for large grid without pattern")
	}
}

func replaceAt(original string, start int, replacement string) string {
	return original[:start] + replacement + original[start+len(replacement):]
}
