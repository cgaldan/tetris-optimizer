package tetro

import "fmt"

// createGrid initializes a square grid of the given size filled with '.' characters.
func createGrid(size int) [][]rune {
	grid := make([][]rune, size)
	for i := range grid {
		grid[i] = make([]rune, size)
		for j := range grid[i] {
			grid[i][j] = '.'
		}
	}
	return grid
}

// PrintGrid prints the grid row by row.
func PrintGrid(grid [][]rune) {
	for _, line := range grid {
		fmt.Println(string(line))
	}
}

// canPlace checks if the tetromino can be placed at position (x, y) on the grid.
// It ensures that tetromino '#' blocks stay within grid bounds and don't overlap existing placements.
func canPlace(grid [][]rune, tetromino [4][4]rune, x, y int) bool {
	gridSize := len(grid)

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if tetromino[i][j] == '#' {
				if x+i >= gridSize || y+j >= gridSize || grid[x+i][y+j] != '.' {
					return false
				}
			}
		}
	}
	return true
}

// placeTetromino places the tetromino on the grid at position (x, y) using the given letter.
func placeTetromino(grid [][]rune, tetromino [4][4]rune, x, y int, letter rune) {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if tetromino[i][j] != '.' {
				grid[x+i][y+j] = letter
			}
		}
	}
}

// removeTetromino removes the tetromino from the grid at position (x, y),
// restoring the cells back to the '.' character.
func removeTetromino(grid [][]rune, tetromino [4][4]rune, x, y int) {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if tetromino[i][j] != '.' {
				grid[x+i][y+j] = '.'
			}
		}
	}
}
