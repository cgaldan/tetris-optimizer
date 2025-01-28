package main

import (
	"fmt"
	"os"
	tetro "tetris/utils"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <filename>")
		return
	}

	filename := os.Args[1]
	tetrominoes, err := tetro.ReadTetrominoes(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	for i, t := range tetrominoes {
		fmt.Printf("Tetromino %d:\n", i+1)
		for _, line := range t {
			fmt.Println(string(line[:]))
		}
		fmt.Println()
	}

	tetrominoes = tetro.PreprocessTetrominoes(tetrominoes)
	grid := tetro.FindSmallestGrid(tetrominoes)
	tetro.PrintGrid(grid)
}
