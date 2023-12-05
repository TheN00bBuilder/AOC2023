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
	handle, error := os.Open("d4i.txt")
	if error != nil {
		log.Fatal(error)
	}
	scanner := bufio.NewScanner(handle)
	num_str := ""
	running_count := 0
	card := false
	winner := false
	card_count := 0
	// for each line
	for scanner.Scan() {
		winnum_map := make(map[int]int)
		card = false
		winner = false
		num_str = ""
		previous := rune('b') //B
		// for each char
		for _, runeval := range scanner.Text() {
			if unicode.IsDigit(runeval) {
				num_str = num_str + string(runeval)
			}
			// if its a Card number, throw out the number and set winner
			if runeval == ':' {
				num_str = ""
				winner = true
			}
			// if space and not empty, add to list
			if runeval == ' ' && card {
				// 2 spaces signify a single num, if last was a space then it isn't a number
				if previous != ' ' {
					myint, _ := strconv.Atoi(num_str)
					// if not populated
					if winnum_map[myint] != 0 {
						// edge case - if 0, make it 1
						if card_count == 0 {
							card_count = 1
						} else {
							// else multiply
							card_count = card_count * 2
						}
					}
				}
				num_str = ""
			}
			// if we hit a space and we're processing the winner card, add it to the map
			if runeval == ' ' && winner {
				// if the previous number was not a space (aka 2 spaces in a row, a single num), this is a single number
				// process it
				if previous != ' ' {
					winint, _ := strconv.Atoi(num_str)
					winnum_map[winint] = winint
				}
				num_str = ""
			}
			// if we hit a bar, switch card types
			if runeval == '|' {
				winner = false
				card = true
			}
			previous = runeval
		}
		// Check last number, wow this method SUCKS
		myint, _ := strconv.Atoi(num_str)
		if winnum_map[myint] != 0 {
			if card_count == 0 {
				card_count = 1
			} else {
				card_count = card_count * 2
			}
		}
		running_count = running_count + card_count
		card_count = 0

	}
	fmt.Println(running_count)
}
