# **Tetris Optimizer**

## **Description**
This Program reads tetrominoes from a text file, finds the optimal placement to fit them into the smallest square grid possible (minimizing empty spaces), and displays the solution board with colorful tetromino blocks. It also shows each tetromino individually and includes a live timer that displays the elapsed time while finding the solution.



## **File Format**

**The input file should contain tetrominoes in a specific 4x4 grid format:**
- Each tetromino is represented as a 4x4 grid using only `#` (block) and `.` (empty space).
- Tetrominoes are separated by one or more blank lines.

**Example:**
    
    ....
    ..##
    .##.
    ....

    ....
    .##.
    .##.
    ....

## **Run the Project**
**To run this program golang is needed to be installed**
1. **Clone the repository:**
   ```bash
   git clone https://platform.zone01.gr/git/cgkaldan/tetris-optimizer.git
   cd tetris-optimizer
2. **Run the program:**
- In "filename" argument the user must use a .txt file from the examples folder
    ```bash 
    go run main.go "filename"
    ```

## **Project Structure**

1. **examples:**
    - In the examples folder there are some ready test case that the user may use to run the program
2. **frontend:**
    - In frontend folder the user interface elements are being handled like displaying the grid, tetromino previews and timer.
3. **utils:**
    - In utils folder there is the core logic for parsing, validating and solving tetromino placements.
    

## **Features**

1. **Tetromino Parsing:** Reads and validates tetromino shapes from a text file.
2. **Live Timer:** Shows a live countdown of the time taken to find the solution.
3. **Optimal Placement:** Uses a recursive algorithm to fit tetrominoes in the smallest square grid.
4. **Tetromino Preview:** Displays each tetromino separately in the same styled format as the solution board.
5. **Colorful Output:** Prints the solution board with colors that make the tetromino blocks visually appealing.

## **Authors**
 -  Christos Gkaldanidis
 
 Creator and primary Developer