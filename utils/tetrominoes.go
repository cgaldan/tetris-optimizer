package tetro

import (
	"bufio"
	"os"
	"strings"
)

// ReadTetrominoes reads a file containing tetromino shapes and returns a slice of 4x4 rune arrays.
func ReadTetrominoes(filename string) [][4][4]rune {
	filename = "examples/" + filename

	file, err := os.Open(filename)
	if err != nil {
		return nil
	}
	defer file.Close()

	var tetrominoes [][4][4]rune
	var current []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.TrimSpace(line) == "" {
			// Process the collected lines when an empty line is encountered
			if len(current) > 0 {
				tetromino := formatCheck(current)
				if tetromino == [4][4]rune{} {
					return nil // Invalid tetromino format
				}
				tetrominoes = append(tetrominoes, tetromino)
				current = nil
			}
			continue
		}

		current = append(current, line)
	}

	if err := scanner.Err(); err != nil {
		return nil
	}

	// Process the last tetromino if the file doesn't end with an empty line
	if len(current) > 0 {
		tetromino := formatCheck(current)
		if tetromino == [4][4]rune{} {
			return nil
		}
		tetrominoes = append(tetrominoes, tetromino)
	}

	return tetrominoes
}

// formatCheck validates and converts a slice of strings into a 4x4 rune array.
func formatCheck(lines []string) [4][4]rune {
	var tetromino [4][4]rune

	if len(lines) != 4 {
		return [4][4]rune{} // Ensure tetromino has exactly 4 rows
	}

	for i, line := range lines {
		if len(line) != 4 {
			return [4][4]rune{} // Ensure each row has exactly 4 columns
		}

		for j, char := range line {
			if char != '.' && char != '#' {
				return [4][4]rune{} // Only '.' and '#' are allowed
			}
			tetromino[i][j] = char
		}
	}

	if !tetrominoCheck(tetromino) {
		return [4][4]rune{} // Ensure the tetromino is valid and connected
	}

	return tetromino
}

// tetrominoCheck verifies that the tetromino consists of exactly 4 connected '#' blocks.
func tetrominoCheck(tetromino [4][4]rune) bool {
	visited := [4][4]bool{}
	var blocks [][2]int

	// Collect all '#' block positions
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if tetromino[i][j] == '#' {
				blocks = append(blocks, [2]int{i, j})
			}
		}
	}

	if len(blocks) != 4 {
		return false // Tetromino must have exactly 4 blocks
	}

	var queue [][2]int
	queue = append(queue, blocks[0])
	visited[blocks[0][0]][blocks[0][1]] = true

	count := 0
	directions := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} // Right, Left, Down, Up

	// BFS to check if all 4 '#' are connected
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		count++

		for _, dir := range directions {
			ni, nj := current[0]+dir[0], current[1]+dir[1]

			if ni >= 0 && ni < 4 && nj >= 0 && nj < 4 && tetromino[ni][nj] == '#' && !visited[ni][nj] {
				visited[ni][nj] = true
				queue = append(queue, [2]int{ni, nj})
			}
		}
	}

	return count == 4 // Valid if all 4 blocks are connected
}

// trimTetromino removes empty spaces "." around the tetromino.
func trimTetromino(tetromino [4][4]rune) [4][4]rune {
	var minRow, maxRow, minCol, maxCol int
	minRow, minCol = 4, 4
	maxRow, maxCol = -1, -1

	// Find the bounding box of the tetromino
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if tetromino[i][j] == '#' {
				if i < minRow {
					minRow = i
				}
				if i > maxRow {
					maxRow = i
				}
				if j < minCol {
					minCol = j
				}
				if j > maxCol {
					maxCol = j
				}
			}
		}
	}

	// Create a new trimmed tetromino
	var trimmed [4][4]rune
	for i := range trimmed {
		for j := range trimmed[i] {
			trimmed[i][j] = '.'
		}
	}

	// Copy only the relevant part
	for i := minRow; i <= maxRow; i++ {
		for j := minCol; j <= maxCol; j++ {
			trimmed[i-minRow][j-minCol] = tetromino[i][j]
		}
	}

	return trimmed
}

// PreprocessTetrominoes trims all tetrominoes to their bounding box.
func PreprocessTetrominoes(tetrominoes [][4][4]rune) [][4][4]rune {
	trimmedTetrominoes := make([][4][4]rune, len(tetrominoes))
	for i, tetromino := range tetrominoes {
		trimmedTetrominoes[i] = trimTetromino(tetromino)
	}
	return trimmedTetrominoes
}
