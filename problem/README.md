# Grid Problem Implementation

## Strategy

1. Use the first pattern row to find candidate starting positions.

2. For each candidate, verify the remaining pattern rows at the same column.

> This avoids checking every cell blindly and keeps the implementation simple, readable, and efficient.

## Complexity

Let:

- $R$ = number of rows in the grid
- $C$ = average number of columns per grid row
- $P$ = number of rows in the pattern
- $W$ = Width of each pattern row
- $K$ = number of times the first pattern row appears as a candidate in the grid

The algorithm works as follows:

```text
Grid:
R rows × C columns

Pattern:
P rows × W columns
```

First it searches for the first pattern row:

```text
Find this:
1920

Inside grid rows:
12930482091238127
12037128741023801
19208301741029388  <- candidate row
19238012938172401
19123812038127415
```

Then validates downwards from that candidate row:

```text
Candidate at row 2, col 0:

1920  matches pattern row 0
1923  matches pattern row 1
1912  matches pattern row 2

Pattern found.
```

The practical time complexity is $$O((R - P + 1) * C + K * P * W)$$

Meaning:

1. Scan possible starting rows
2. Find candidate colums using `strings.Index`
3. For each candidate, check the next P-1 rows for a match

Worst case is when the first pattern row appears frequently ($O(R * C * P * W)$), but in practice this is efficient for typical grid and pattern sizes.

Space complexity is $$O(1)$$ since we only use a few variables for tracking positions and matches, and do not store any additional data structures proportional to the input size.

### Issues or Challenges

The main challenge is avoiding unnecessary comparisons. A brute-force implementation would check every possible row and column combination. This implementation reduces unnecessary work by first searching for the first pattern row, then only validating full pattern matches from those candidate positions.

### Extras

To make this usable by other applications, I would expose the solver behind a small package-level API, document the expected input behavior, and add examples. For larger-scale use cases, I would consider adding input validation errors instead of returning only a boolean, plus benchmarks for large grids.

Improvements could include:

To make this usable by other applications, I would expose the solver behind a small package-level API, document the expected input behavior, and add examples. For larger-scale use cases, I would consider adding input validation errors instead of returning only a boolean, plus benchmarks for large grids.

Improvements could include:

- Return match coordinates instead of only true/false
- Add an error-returning API for invalid inputs
- Add benchmark coverage for different grid sizes
- Support streaming input for very large grids
- Consider more advanced string matching if the first pattern row appears very frequently
