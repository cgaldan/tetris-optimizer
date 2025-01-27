package tetro

import (
	"math"
)

func Solution(grid [][]rune, tetrominoes [][4][4]rune, index int) bool {
	if index == len(tetrominoes) {
		return true
	}

	tetromino := tetrominoes[index]
	gridSize := len(grid)
	letter := 'A' + rune(index)

	for x := 0; x < gridSize; x++ {
		for y := 0; y < gridSize; y++ {
			if canPlace(grid, tetromino, x, y) {
				placeTetromino(grid, tetromino, x, y, letter)
				if Solution(grid, tetrominoes, index+1) {
					return true
				}
				removeTetromino(grid, tetromino, x, y)
			}
		}
	}
	// PrintGrid(grid)
	return false
}

func FindSmallestGrid(tetrominoes [][4][4]rune) [][]rune {
	size := minGridSize(tetrominoes)

	for {
		grid := createGrid(size)
		// fmt.Println(size)
		if Solution(grid, tetrominoes, 0) {

			return grid
		}
		// PrintGrid(grid)
		size++
	}
}

func minGridSize(tetrominoes [][4][4]rune) int {
	totalBlocks := len(tetrominoes) * 4
	size := int(math.Ceil(math.Sqrt(float64(totalBlocks))))

	return size
}

// func FindOptimalGrid(tetrominoes [][4][4]rune) int {
// 	minSize := int(math.Ceil(math.Sqrt(float64(len(tetrominoes) * 4))))
// 	maxSize := minSize * 2

// 	for minSize <= maxSize {
// 		midSize := (minSize + maxSize) / 2
// 		grid := CreateGrid(midSize)

// 		if Solution(grid, tetrominoes, 0) {
// 			minSize = midSize
// 		} else {
// 			maxSize = midSize + 1
// 		}
// 	}

// 	return minSize
// }
