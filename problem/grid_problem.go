package problem

// IGridProblem defines the interface for solving the grid pattern matching problem.
//
// Given a matrix of characters, check to see if the pattern grid exists inside of the larger grid.
//
// Grid example:
//
//	12930482091238127
//	12037128741023801
//	19208301741029388
//	19238012938172401
//	19123812038127415
//
// grid[0] = first row, grid[1] = second row, ..., grid[n] = nth row
//
// Pattern example:
//
//	1920
//	1923
//	1912
//
// solveGrid returns true if the pattern exists within the main grid, false otherwise.
type IGridProblem interface {
	SolveGrid(grid []string, pattern []string) bool
}
