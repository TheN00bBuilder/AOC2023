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
	line_count := 1
	copy_map := make(map[int]int)
	for i := 1; i < 217; i++ {
		copy_map[i] = 1
	}
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
					if winnum_map[myint] != 0 {
						//winner
						card_count++
					}
				}
				num_str = ""
			}
			// if end of number and we're processing a winner card
			if runeval == ' ' && winner {
				if previous != ' ' {
					winint, _ := strconv.Atoi(num_str)
					winnum_map[winint] = winint
				}
				num_str = ""
			}
			if runeval == '|' {
				winner = false
				card = true
			}
			previous = runeval
		}
		// Check last number, wow this method SUCKS
		myint, _ := strconv.Atoi(num_str)
		if winnum_map[myint] != 0 {
			card_count++
		}
		// store line count of current card
		ihateeverything := copy_map[line_count]
		// for each of the current cards
		for i := 0; i < ihateeverything; i++ {
			// for each of the cards that I won copies of
			for j := 0; j < card_count; j++ {
				// add 1 to each of the cards i won
				copy_map[line_count+j+1] = copy_map[line_count+j+1] + 1
			}
		}
		// line goes up
		line_count++
		// cards go reset
		card_count = 0
	}
	for i := 1; i < len(copy_map)+1; i++ {
		// count card total
		running_count = running_count + copy_map[i]
	}
	fmt.Println(running_count)
}
