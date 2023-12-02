package main

// don't judge a man based off his imports
import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	handle, error := os.Open("d2i.txt")
	if error != nil {
		log.Fatal(error)
	}
	scanner := bufio.NewScanner(handle)
	running_count := 0
	// yay less regex
	redregex, _ := regexp.Compile("[0-9]{1,2} red")
	greenregex, _ := regexp.Compile("[0-9]{1,2} green")
	blueregex, _ := regexp.Compile("[0-9]{1,2} blue")
	for scanner.Scan() {
		line := scanner.Text()
		redcount := 0
		bluecount := 0
		greencount := 0
		semicolon := strings.Split(line, ";")
		// for each set in each game
		for _, line := range semicolon {
			// find all entries of each color
			for _, element := range redregex.FindAllString(line, -1) {
				// get its number
				observed_str := element[:2]
				number, _ := strconv.Atoi(strings.TrimSpace(observed_str))
				// if it's lower than the current number for the color, replace it
				if number > redcount {
					redcount = number
				}
				observed_str = ""
			}
			for _, element := range greenregex.FindAllString(line, -1) {
				observed_str := element[:2]
				number, _ := strconv.Atoi(strings.TrimSpace(observed_str))
				if number > greencount {
					greencount = number
				}
				observed_str = ""
			}
			for _, element := range blueregex.FindAllString(line, -1) {
				observed_str := element[:2]
				number, _ := strconv.Atoi(strings.TrimSpace(observed_str))
				if number > bluecount {
					bluecount = number
				}
				observed_str = ""
			}
		}
		// instead of summing game IDs, just find the power
		running_count = running_count + (bluecount * redcount * greencount)
		redcount = 0
		bluecount = 0
		greencount = 0
	}
	fmt.Println(running_count)
}
