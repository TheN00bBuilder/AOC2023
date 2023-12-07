package main

// don't judge a man based off his imports
import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	// sorry about hardcoding it, its just so short so what's the point
	inputstr := "put your input here, with lines separated by a space"
	wincount := 0
	// remove ALL things not numbers
	remove_unwanted := regexp.MustCompile("[^a-zA-z ,]")
	modinputstr := remove_unwanted.FindAllString(inputstr, -1)
	// then split off colons
	parts := strings.Split(strings.Join(modinputstr, ""), ":")
	// get time, distance
	timeint, _ := strconv.Atoi(parts[1])
	distint, _ := strconv.Atoi(parts[2])
	for j := 1; j < timeint; j++ {
		// speed is j meters per second
		// but also how long you hold down button
		// so just do the basic math
		distance_travelled := (timeint - j) * j
		if distance_travelled > distint {
			wincount++
		}
	}
	fmt.Println(wincount)
}
