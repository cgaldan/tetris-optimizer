package main

import (
	"fmt"
	"os"
	ui "tetris/frontEnd"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <filename>")
		return
	}
	filename := os.Args[1]

	ui.DisplayWelcomeMessage()
	ui.WholeProcess(filename)
}
