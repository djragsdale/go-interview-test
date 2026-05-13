package problem

import "strings"

// GridProblemImpl implements IGridProblem.
type GridProblemImpl struct{}

// SolveGrid provides the implementation to solve the problem statement.
func (g GridProblemImpl) SolveGrid(grid []string, pattern []string) bool {
	// TODO: Implement the solution
	patternRows := len(pattern)

	// Empty pattern trivially found
	if patternRows == 0 {
		return true
	}

	gridRows := len(grid)

	// If the grid has fewer rows than the pattern, it cannot contain the pattern
	if gridRows == 0 || patternRows > gridRows {
		return false
	}

	patternWidth := len(pattern[0])

	// Empty pattern rows trivially found as long as there are enough rows
	if patternWidth == 0 {
		return true
	}

	// Validate that the pattern is rectangular
	for _, row := range pattern {
		if len(row) != patternWidth {
			return false // Non-rectangular pattern
		}
	}

	firstPatternRow := pattern[0]

	// Only search rows where the full patterc can fir
	for gridRow := 0; gridRow <= gridRows-patternRows; gridRow++ {
		row := grid[gridRow]

		if len(row) < patternWidth {
			continue // Skip rows that are too short to contain the pattern
		}

		searchStart := 0

		// Find every ocurrence of the first pattern row in the grid row
		for searchStart <= len(row)-patternWidth {
			idx := strings.Index(row[searchStart:], firstPatternRow)
			if idx == -1 {
				break
			}
			gridCol := searchStart + idx

			if matchesPattern(grid, pattern, gridRow, gridCol, patternWidth) {
				return true
			}
			searchStart = gridCol + 1 // Move past the current match
		}
	}
	return false
}

// matchesPattern checks if the pattern matches the grid starting at the specified position.
func matchesPattern(grid []string, pattern []string, startRow, startCol, patternWidth int) bool {
	// pattern[0] is already matched, so we start checking from the second row of the pattern
	for patternRow := 1; patternRow < len(pattern); patternRow++ {
		gridRow := grid[startRow+patternRow]

		if startCol+patternWidth > len(gridRow) {
			return false // Grid row is too short to match the pattern
		}

		if gridRow[startCol:startCol+patternWidth] != pattern[patternRow] {
			return false // Pattern row does not match the grid row
		}
	}
	return true
}
