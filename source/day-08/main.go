package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type display struct {
	signalPatterns []string
	outputValues   []string
}

func main() {
	fmt.Println(part1(readInput("input-08.txt")))
}

func part1(displays []display) int {
	var uniqueDigitCount int
	for i := 0; i < len(displays); i++ {
		for j := 0; j < len(displays[i].outputValues); j++ {
			//len is only appropriate to use to determine number of characters in a string if
			//the string consists solely of byte-sized characters. *chomp* This is luckily the case here.
			length := len(displays[i].outputValues[j])
			//Digits with unique amount of segements are 1 (2 seg), 4 (4 seg), 7 (3 seg) and 8 (7 seg).
			if length == 2 || length == 3 || length == 4 || length == 7 {
				uniqueDigitCount = uniqueDigitCount + 1
			}
		}
	}
	return uniqueDigitCount
}

func readInput(path string) []display {
	f, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var displays []display

	for scanner.Scan() {
		input := strings.Split(scanner.Text(), " | ")
		firstPart := input[0]
		secondPart := input[1]
		display := display{}
		display.signalPatterns = strings.Split(firstPart, " ")
		display.outputValues = strings.Split(secondPart, " ")
		displays = append(displays, display)
	}

	return displays
}
