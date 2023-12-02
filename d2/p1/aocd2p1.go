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
	gamenum := ""
	running_count := 0
	my_id := 0
	// I HATE REGEX
	gameregex, _ := regexp.Compile("[^A-Za-z\\s][0-9]*")
	redregex, _ := regexp.Compile("[0-9]{1,2} red")
	greenregex, _ := regexp.Compile("[0-9]{1,2} green")
	blueregex, _ := regexp.Compile("[0-9]{1,2} blue")
	for scanner.Scan() {
		line := scanner.Text()
		my_id = 0
		redcount := 0
		bluecount := 0
		greencount := 0
		gamenum = gameregex.FindAllString(line[:8], -1)[0]
		my_id, _ = strconv.Atoi(gamenum)
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
		// if the highest number in the set satisfies the issue, add the ID to the running count
		if redcount <= 12 && greencount <= 13 && bluecount <= 14 {
			running_count = running_count + my_id
		}
		redcount = 0
		bluecount = 0
		greencount = 0
	}
	fmt.Println(running_count)
}
