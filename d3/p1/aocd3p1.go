package main

// don't judge a man based off his imports
import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

const ROWS = 140
const COLUMNS = 140

func main() {
	handle, error := os.Open("d3i.txt")
	if error != nil {
		log.Fatal(error)
	}
	scanner := bufio.NewScanner(handle)
	twodim_array := [ROWS][COLUMNS]rune{}
	linecount := 0
	// for each line
	for scanner.Scan() {
		// for each char
		for count, runeval := range scanner.Text() {
			twodim_array[linecount][count] = runeval
		}
		linecount++
	}
	current_num := ""
	running_count := 0
	hasSurrounded := false
	// 2d array of each line now exists
	for i := 0; i < ROWS; i++ {
		for j := 0; j < COLUMNS; j++ {
			// iterate over 2D array, if its a digit, its under investigation
			if unicode.IsDigit(twodim_array[i][j]) {
				current_num = current_num + string(twodim_array[i][j])
				// if surrounded by something, set the flag
				if isSurrounded(twodim_array, i, j) {
					hasSurrounded = true
				}
			}
			if !(unicode.IsDigit(twodim_array[i][j])) {
				// number finished - if we've processed the number, and it is surrounded by anything, add to total
				if hasSurrounded {
					intval, _ := strconv.Atoi(current_num)
					running_count = running_count + intval
					hasSurrounded = false
				}
				// reset number
				current_num = ""
			}
		}
	}
	fmt.Println(running_count)
}

/*

v v v * *
v 1 v 3 *
v v v * *

possibilities for [r][c]
[r-1][c-1] (top left) v
[r-1][c] (above) v
[r-1][c+1] (top right) v
[r][c-1] (immediate left)
[r][c+1] (immediate right)
[r+1][c-1] (bottom left)
[r+1][c] (below)
[r+1][c+1] (bottom right)



*/

func isSurrounded(arr [ROWS][COLUMNS]rune, row int, col int) bool {
	// if top left exists
	if (row-1 != -1) && (col-1 != -1) {
		if !unicode.IsDigit(arr[row-1][col-1]) && arr[row-1][col-1] != '.' {
			return true
		}
	}
	if row-1 != -1 {
		// follows that row-1, col will exist
		if !unicode.IsDigit(arr[row-1][col]) && arr[row-1][col] != '.' {
			return true
		}
	}
	// now check if row-1 exists
	if col-1 != -1 {
		if !unicode.IsDigit(arr[row][col-1]) && arr[row][col-1] != '.' {
			return true
		}
	}
	// now check if top right exists
	if (row-1 != -1) && (col+1 != COLUMNS) {
		if !unicode.IsDigit(arr[row-1][col+1]) && arr[row-1][col+1] != '.' {
			return true
		}
	}
	// now check if r+1, col-1 exists
	if (row+1 != ROWS) && (col-1 != -1) {
		if !unicode.IsDigit(arr[row+1][col-1]) && arr[row+1][col-1] != '.' {
			return true
		}
	}
	// now check if r+1, col exists
	if row+1 != ROWS {
		if !unicode.IsDigit(arr[row+1][col]) && arr[row+1][col] != '.' {
			return true
		}
	}

	// now check if r+1, c+1 exists
	if (row+1 != ROWS) && (col+1 != COLUMNS) {
		if !unicode.IsDigit(arr[row+1][col+1]) && arr[row+1][col+1] != '.' {
			return true
		}
	}

	if col+1 != COLUMNS {
		if !unicode.IsDigit(arr[row][col+1]) && arr[row][col+1] != '.' {
			return true
		}
	}
	return false
}
