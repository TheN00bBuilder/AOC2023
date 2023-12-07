package main

// don't judge a man based off his imports
import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const ELEMENTS = 4

func main() {
	// sorry about hardcoding it, its just so short so what's the point
	inputstr := "put your input here, with your lines separated by a newline (\\n)"
	// remove the whitespace
	remove_whitespace := regexp.MustCompile(`\s+`)
	inputstr = remove_whitespace.ReplaceAllString(inputstr, " ")
	parts := strings.Split(inputstr, " ")
	wincount := 0
	multi_answer := 1
	for i := 1; i <= ELEMENTS; i++ {
		timeint, _ := strconv.Atoi(parts[i])
		distint, _ := strconv.Atoi(parts[i+ELEMENTS+1])
		for j := 1; j < timeint; j++ {
			// speed is j meters per second
			// but also how long you hold down button
			// so just do the basic math
			distance_travelled := (timeint - j) * j
			if distance_travelled > distint {
				wincount++
			}
		}
		// multiply for answer, reset wincount
		multi_answer = multi_answer * wincount
		wincount = 0
	}
	fmt.Println(multi_answer)
}
