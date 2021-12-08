package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type display struct {
	signalPatterns []string
	outputValues   []string
}

func main() {
	readInput("input-08-example.txt")
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
