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

// WholeProcess handles the complete process from reading tetrominoes to displaying the solution.
func WholeProcess(filename string) {
	startTime := time.Now()
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	// Channel to signal when the process is complete.
	done := make(chan bool)
	go func() {

		tetrominoes := tetro.ReadTetrominoes(filename)
		if tetrominoes == nil {
			fmt.Println(red + "ERROR\n" + reset)
			done <- true
			return
		}

		tetrominoes = tetro.PreprocessTetrominoes(tetrominoes)
		grid := tetro.FindSmallestGrid(tetrominoes)
		fmt.Print("\r                    ")

		// Print each tetromino with its corresponding letter.
		x, y := 0, 0
		for letter, tetromino := range tetrominoes {
			PrintTetromino(tetromino, letter)
			// Arrange tetromino previews in rows.
			y += 4
			if y >= len(grid[0]) {
				y = 0
				x += 4
			}
		}

		DisplayStatistics(len(grid), startTime, len(tetrominoes))
		PrintGridWithBorders(grid)
		fmt.Println()
		done <- true
	}()

	// Timer display until the process is finished.
	for {
		select {
		case <-done:
			return // Exit once the process is done.
		case <-ticker.C:
			sec := time.Since(startTime).Seconds()
			// Use carriage return to update timer on the same line.
			fmt.Printf(gray+"\rTimer: %.2f seconds"+reset, sec)

		}
	}
}

// DisplayWelcomeMessage shows a welcome screen and waits for user input before proceeding.
func DisplayWelcomeMessage() {
	fmt.Println(clear)
	fmt.Println(magenta + "\n" + `╔╦╗╔═╗╔╦╗╦═╗╦╔═╗  ╔═╗╔═╗╔╦╗╦╔╦╗╦╔═╗╔═╗╦═╗
 ║ ║╣  ║ ╠╦╝║╚═╗  ║ ║╠═╝ ║ ║║║║║╔═╝║╣ ╠╦╝
 ╩ ╚═╝ ╩ ╩╚═╩╚═╝  ╚═╝╩   ╩ ╩╩ ╩╩╚═╝╚═╝╩╚═` + reset)
	fmt.Println()
	fmt.Println(gray + "Press " + reset + yellow + "Enter " + gray + "to start solving..." + reset)
	fmt.Scanln() // Wait for the user to press Enter.
}

// DisplayStatistics shows the solution statistics including tetromino count, grid size and time taken.
func DisplayStatistics(gridSize int, startTime time.Time, tetrominoes int) {
	elapsed := time.Since(startTime).Seconds()
	fmt.Printf(clear + green + "\n\n✓Solution Found!\n\n" + reset)
	fmt.Printf(gray+"Amount of tetrominoes: "+reset+blue+"%d\n"+reset, tetrominoes)
	fmt.Printf(gray+"Smallest grid size: "+reset+blue+"%dx%d\n"+reset, gridSize, gridSize)
	fmt.Printf(gray+"Time taken: "+reset+blue+"%.2f seconds\n\n"+reset, elapsed)
}

// GetTetrominoColor returns a color code based on the tetromino letter.
func GetTetrominoColor(letter rune) string {
	// Define a color palette.
	colors := []string{
		"\033[41m",               // Red
		"\033[42m",               // Green
		"\033[44m",               // Blue
		"\033[43m",               // Orange
		"\033[45m",               // Magenta
		"\033[46m",               // Cyan
		"\033[48;5;226m",         // Yellow
		"\033[48;2;0;255;127m",   // Bright Green
		"\033[48;2;255;105;180m", // Pink
		"\033[48;2;153;0;76m",    // Purple
		"\033[48;2;51;0;51m",     // Dark Purple
		"\033[48;2;80;0;0m",      // Dark Red
	}

	// Use modulo to cycle through available colors.
	colorIndex := (int(letter) - 'A') % len(colors)
	return colors[colorIndex]
}

// PrintGridWithBorders prints the grid with borders and row/column numbers.
func PrintGridWithBorders(grid [][]rune) {
	gridSize := len(grid)

	// Print top border with column numbers.
	fmt.Print("    ")
	for i := 1; i <= gridSize; i++ {
		fmt.Printf(" %d ", i)
	}
	fmt.Println()
	fmt.Print("   ┌")
	for i := 0; i < gridSize; i++ {
		fmt.Print("───")
	}
	fmt.Println("┐")

	// Print each row with side borders.
	for i, row := range grid {
		fmt.Printf("%2d │", i+1) // Row number.
		for _, cell := range row {
			if cell == '.' {
				fmt.Print(" • ") // Empty cell indicator.
			} else {
				// Color the tetromino block based on its letter.
				color := GetTetrominoColor(cell)             // background color
				color1 := fmt.Sprintf("\033[3%s", color[3:]) // text color
				fmt.Printf("%s%s . \033[0m", color, color1)
			}
		}
		fmt.Println("│")
	}

	// Print bottom border.
	fmt.Print("   └")
	for i := 0; i < gridSize; i++ {
		fmt.Print("───")
	}
	fmt.Println("┘")
}

// PrintTetromino prints an individual tetromino with its letter and color.
func PrintTetromino(tetromino [4][4]rune, index int) {
	letter := 'A' + rune(index)
	color := GetTetrominoColor(letter)           //background color
	color1 := fmt.Sprintf("\033[3%s", color[3:]) // text color
	fmt.Printf(color1+"\nTetromino %c:\n\n"+reset, letter)
	for _, row := range tetromino {
		for _, cell := range row {
			if cell == '#' {
				fmt.Printf("%s%s . \033[0m", color, color1) // Display tetromino block.
			} else {
				fmt.Print("   ") // Display empty space.
			}
		}
		fmt.Println()
	}
}
