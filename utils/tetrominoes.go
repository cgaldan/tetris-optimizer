package tetro

import (
	"bufio"
	"os"
	"strings"
)

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
			if len(current) > 0 {
				tetromino := formatCheck(current)
				if tetromino == [4][4]rune{} {
					return nil
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

	if len(current) > 0 {
		tetromino := formatCheck(current)
		if tetromino == [4][4]rune{} {
			return nil
		}
		tetrominoes = append(tetrominoes, tetromino)
	}

	return tetrominoes
}

func formatCheck(lines []string) [4][4]rune {
	var tetromino [4][4]rune

	if len(lines) != 4 {
		return [4][4]rune{}
	}

	for i, line := range lines {
		if len(line) != 4 {
			return [4][4]rune{}
		}

		for j, char := range line {
			if char != '.' && char != '#' {
				return [4][4]rune{}
			}
			tetromino[i][j] = char
		}
	}

	if !tetrominoCheck(tetromino) {
		return [4][4]rune{}
	}

	return tetromino
}

func tetrominoCheck(tetromino [4][4]rune) bool {
	visited := [4][4]bool{}
	var blocks [][2]int

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if tetromino[i][j] == '#' {
				blocks = append(blocks, [2]int{i, j})
			}
		}
	}

	if len(blocks) != 4 {
		return false
	}

	var queue [][2]int
	queue = append(queue, blocks[0])
	visited[blocks[0][0]][blocks[0][1]] = true

	count := 0
	directions := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

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

	return count == 4
}

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

	for i := minRow; i <= maxRow; i++ {
		for j := minCol; j <= maxCol; j++ {
			trimmed[i-minRow][j-minCol] = tetromino[i][j]
		}
	}

	return trimmed
}

func PreprocessTetrominoes(tetrominoes [][4][4]rune) [][4][4]rune {
	trimmedTetrominoes := make([][4][4]rune, len(tetrominoes))
	for i, tetromino := range tetrominoes {
		trimmedTetrominoes[i] = trimTetromino(tetromino)
	}
	return trimmedTetrominoes
}
