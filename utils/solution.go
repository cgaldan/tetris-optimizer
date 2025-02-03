package tetro

import (
	"math"
)

// Solution recursively places tetrominoes on the grid starting from the given index.
// It returns true if a valid placement is found for all tetrominoes.
func Solution(grid [][]rune, tetrominoes [][4][4]rune, index int) bool {
	if index == len(tetrominoes) {
		return true // All tetrominoes have been placed successfully.
	}

	tetromino := tetrominoes[index]
	gridSize := len(grid)
	letter := 'A' + rune(index) // Use a unique letter for each tetromino.

	// Try to place the current tetromino in every possible grid position.
	for x := 0; x < gridSize; x++ {
		for y := 0; y < gridSize; y++ {
			if canPlace(grid, tetromino, x, y) {
				placeTetromino(grid, tetromino, x, y, letter)
				if Solution(grid, tetrominoes, index+1) {
					return true // Found a valid placement for all tetrominoes.
				}
				removeTetromino(grid, tetromino, x, y) // Backtrack if placement did not lead to a solution.
			}
		}
	}
	return false // No valid placement found for the current configuration.
}

// FindSmallestGrid finds the smallest grid that can fit all tetrominoes.
func FindSmallestGrid(tetrominoes [][4][4]rune) [][]rune {
	size := minGridSize(tetrominoes)
	// Increment grid size until a valid solution is found.
	for {
		grid := createGrid(size)
		if Solution(grid, tetrominoes, 0) {
			return grid
		}
		size++
	}
}

// minGridSize estimates the minimum grid size needed by taking the square root
// of the total number of tetromino blocks and rounding up.
func minGridSize(tetrominoes [][4][4]rune) int {
	totalBlocks := len(tetrominoes) * 4
	size := int(math.Ceil(math.Sqrt(float64(totalBlocks))))
	return size
}
