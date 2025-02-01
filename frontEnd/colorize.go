package ui

import (
	"fmt"
	"time"

	tetro "tetris/utils"
)

const (
	reset   = "\033[0m"
	red     = "\033[31m"
	green   = "\033[32m"
	yellow  = "\033[33m"
	blue    = "\033[34m"
	magenta = "\033[35m"
	gray    = "\033[90m"
	clear   = "\033[2J\033[H"
)

func WholeProcess(filename string) {
	startTime := time.Now()
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	// This channel simulates when your process stops.
	// Replace it with your own termination logic.
	done := make(chan bool)
	go func() {
		tetrominoes := tetro.ReadTetrominoes(filename)
		if tetrominoes == nil {
			fmt.Println(red + "ERROR\n" + reset)
			return
		}

		tetrominoes = tetro.PreprocessTetrominoes(tetrominoes)
		grid := tetro.FindSmallestGrid(tetrominoes)

		x, y := 0, 0
		for letter, tetromino := range tetrominoes {
			PrintTetromino(tetromino, letter)
			// Move to the next position
			y += 4
			if y >= len(grid[0]) {
				y = 0
				x += 4
			}
		}

		DisplayStatistics(len(grid), startTime)
		PrintGridWithBorders(grid)
		fmt.Println()
		done <- true
	}()

	for {
		select {
		case <-done:
			return // Exit if the process is done
		case <-ticker.C:
			// Calculate elapsed time and print it on the same line.
			sec := time.Since(startTime).Seconds()
			// The \r returns the cursor to the start of the line.
			fmt.Printf(gray+"\rTimer: %.2f seconds"+reset, sec)
		}
	}
}

func DisplayWelcomeMessage() {
	fmt.Println(clear)
	fmt.Println(magenta + "\n" + `╔╦╗┌─┐┌┬┐┬─┐┬┌─┐  ╔═╗┌─┐┌┬┐┬┌┬┐┬┌─┐┌─┐┬─┐
 ║ ├┤  │ ├┬┘│└─┐  ║ ║├─┘ │ │││││┌─┘├┤ ├┬┘
 ╩ └─┘ ┴ ┴└─┴└─┘  ╚═╝┴   ┴ ┴┴ ┴┴└─┘└─┘┴└─` + reset)
	fmt.Println()
	fmt.Println(gray + "Press " + reset + yellow + "Enter " + gray + "to start solving..." + reset)
	fmt.Scanln() // Wait for user input before starting
}

func DisplayStatistics(gridSize int, startTime time.Time) {
	elapsed := time.Since(startTime).Seconds()
	fmt.Printf(clear + green + "\n✓Solution Found!\n\n" + reset)
	fmt.Printf(gray+"Smallest grid size: "+reset+blue+"%dx%d\n"+reset, gridSize, gridSize)
	fmt.Printf(gray+"Time taken: "+reset+blue+"%.2f seconds\n\n"+reset, elapsed)
}

func GetTetrominoColor(letter rune) string {
	// Define a color palette
	colors := []string{
		"\033[44m", // Blue
		"\033[42m", // Green
		"\033[41m", // Red
		"\033[43m", // Yellow
		"\033[45m", // Magenta
		"\033[46m", // Cyan
	}

	// Use the modulo operator to cycle through colors
	colorIndex := (int(letter) - 'A') % len(colors)
	if colorIndex < 0 {
		colorIndex += len(colors)
	}
	return colors[colorIndex]
}

func PrintGridWithBorders(grid [][]rune) {
	gridSize := len(grid)

	// Top border
	fmt.Print("   ")
	for i := 1; i <= gridSize; i++ {
		fmt.Printf(" %d ", i)
	}
	fmt.Println()
	fmt.Print("   ┌")
	for i := 0; i < gridSize; i++ {
		fmt.Print("───")
	}
	fmt.Println("┐")

	// Grid rows with side borders
	for i, row := range grid {
		fmt.Printf("%2d │", i+1) // Row number
		for _, cell := range row {
			if cell == '.' {
				fmt.Print(" • ") // Empty cell
			} else {
				// Dynamically fetch the color for the tetromino
				color := GetTetrominoColor(cell)
				fmt.Printf("%s █ \033[0m", color) // Apply color
			}
		}
		fmt.Println("│") // Right border
	}

	// Bottom border
	fmt.Print("   └")
	for i := 0; i < gridSize; i++ {
		fmt.Print("───")
	}
	fmt.Println("┘")
}

func PrintTetromino(tetromino [4][4]rune, index int) {
	letter := 'A' + rune(index)
	color := GetTetrominoColor(letter)
	ansichar := string(color[3])
	colortxt := fmt.Sprintf("\033[3%sm", ansichar)
	fmt.Printf(colortxt+"\nTetromino %c:\n\n"+reset, letter)
	for _, row := range tetromino {
		for _, cell := range row {
			if cell == '#' {
				fmt.Printf("%s █ \033[0m", color) // Blue block for the tetromino
			} else {
				fmt.Print("   ") // Empty cell
			}
		}
		fmt.Println()
	}
}
