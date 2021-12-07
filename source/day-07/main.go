package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(part1())
	fmt.Println(part2())
}

func part1() int {
	crabs := readInput("input-07.txt")
	sort.Ints(crabs)
	smallestCandidate := crabs[0]
	largestCandidate := crabs[len(crabs)-1]
	currentMinimum := -1
	for i := smallestCandidate; i <= largestCandidate; i++ {
		currentFuelCost := calculateFuelCost(crabs, i, currentMinimum)
		if currentFuelCost > -1 {
			if currentMinimum == -1 {
				currentMinimum = currentFuelCost
			} else if currentMinimum > currentFuelCost {
				currentMinimum = currentFuelCost

			}
		}
	}
	return currentMinimum
}

func part2() int {
	crabs := readInput("input-07.txt")
	sort.Ints(crabs)
	smallestCandidate := crabs[0]
	largestCandidate := crabs[len(crabs)-1]
	currentMinimum := -1
	for i := smallestCandidate; i <= largestCandidate; i++ {
		currentFuelCost := calculateFuelCost2(crabs, i, currentMinimum)
		if currentFuelCost > -1 {
			if currentMinimum == -1 {
				currentMinimum = currentFuelCost
			} else if currentMinimum > currentFuelCost {
				currentMinimum = currentFuelCost

			}
		}
	}
	return currentMinimum
}

func calculateFuelCost(crabs []int, candidate int, currentMinimum int) int {
	var totalFuelCost int
	if currentMinimum == -1 {
		for i := 0; i < len(crabs); i++ {
			totalFuelCost = totalFuelCost + absoluteValue(crabs[i]-candidate)
		}
	} else {
		for i := 0; i < len(crabs); i++ {
			if totalFuelCost < currentMinimum {
				totalFuelCost = totalFuelCost + absoluteValue(crabs[i]-candidate)
			} else {
				return -1
			}
		}
	}
	return totalFuelCost
}

func calculateFuelCost2(crabs []int, candidate int, currentMinimum int) int {
	var totalFuelCost int
	if currentMinimum == -1 {
		for i := 0; i < len(crabs); i++ {
			distance := absoluteValue(crabs[i] - candidate)
			//Divergent series
			totalFuelCost = totalFuelCost + (distance*(distance+1))/2
		}
	} else {
		for i := 0; i < len(crabs); i++ {
			if totalFuelCost < currentMinimum {
				distance := absoluteValue(crabs[i] - candidate)
				totalFuelCost = totalFuelCost + (distance*(distance+1))/2
			} else {
				return -1
			}
		}
	}
	return totalFuelCost
}

func absoluteValue(n int) int {
	if n < 0 {
		return -1 * n
	} else {
		return n
	}
}

func readInput(path string) []int {
	f, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	var input string

	for scanner.Scan() {
		input = input + scanner.Text()
	}

	crabsAsStrings := strings.Split(input, ",")
	var crabs []int

	for i := 0; i < len(crabsAsStrings); i++ {
		crabPosition, err := strconv.Atoi(crabsAsStrings[i])
		if err != nil {
			log.Fatal(err)
		}
		crabs = append(crabs, crabPosition)
	}
	return crabs
}
