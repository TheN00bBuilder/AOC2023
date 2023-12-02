package main

// don't judge a man based off his imports
import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	handle, error := os.Open("d1i.txt")
	if error != nil {
		log.Fatal(error)
	}
	scanner := bufio.NewScanner(handle)
	number_string := ""
	running_count := 0
	number_string_array := [10]string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	for scanner.Scan() {
		line := scanner.Text()
		number_string = ""
		// for each element in the number string array
		for counted, element := range number_string_array {
			// get index
			my_index := strings.Index(line, element)
			// ok ok ok ok ok hear me out
			for my_index != -1 {
				// here's where it gets funky:
				// insert whichever one is found INSIDE the number we're working on currently
				// this prevents it from being re-detected forever in a loop
				// because we don't care about what it looks like, only what's in the string
				line = line[:my_index+1] + strconv.Itoa(counted) + line[my_index+1:]
				// then go onto the next match
				my_index = strings.Index(line, element)
			}
		}
		// then we use the same bit as before
		for _, char := range line {
			if unicode.IsNumber(char) {
				number_string = number_string + string(char)
			}
		}
		number_string = string(number_string[0]) + string(number_string[len(number_string)-1])
		scratchint, _ := strconv.Atoi(number_string)
		running_count = running_count + scratchint
	}
	fmt.Print(running_count)
}
