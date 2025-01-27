package tetro

import "fmt"

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

func PrintGrid(grid [][]rune) {
	for _, line := range grid {
		fmt.Println(string(line))
	}
}

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

func placeTetromino(grid [][]rune, tetromino [4][4]rune, x, y int, letter rune) {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if tetromino[i][j] == '#' {
				grid[x+i][y+j] = letter
			}
		}
	}
}

func removeTetromino(grid [][]rune, tetromino [4][4]rune, x, y int) {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if tetromino[i][j] == '#' {
				grid[x+i][y+j] = '.'
			}
		}
	}
}
