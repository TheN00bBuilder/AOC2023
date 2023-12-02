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
	// yes my filename is hardcoded, i don't care
	handle, error := os.Open("d1i.txt")
	if error != nil {
		log.Fatal(error)
	}
	scanner := bufio.NewScanner(handle)
	number_string := ""
	running_count := 0
	for scanner.Scan() {
		// for each line
		number_string = ""
		for _, char := range scanner.Text() {
			// for each char in line, check if number
			if unicode.IsNumber(char) {
				// if number, concatenate to string
				number_string = number_string + string(char)
			}
		}
		// get the first and last numbers into a string
		number_string = string(number_string[0]) + string(number_string[len(number_string)-1])
		// convert it to an integer
		scratchint, _ := strconv.Atoi(number_string)
		// add and reloop
		running_count = running_count + scratchint
	}
	fmt.Print(running_count)
}
