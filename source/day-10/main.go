package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const leftBraces = "([{<"
const rightBraces = ")]}>"
const leftCurve = "("
const rightCurve = ")"
const leftSquare = "["
const rightSquare = "]"
const leftCurly = "{"
const rightCurly = "}"
const leftDiamond = "<"
const rightDiamond = ">"

func main() {
	part1()
}

func part1() {
	lines := readInput("input-10.txt")
	//A left parenthesis can't be illegal; only the right ones.
	//Approach: We find the first right parenthesis. For the line to be legal the one to the left of
	//it has to be its mirror, if it isn't, calculate penalty. If it is, remove both, and check again.
	var totalPenalty int
	for i := 0; i < len(lines); i++ {
		totalPenalty = totalPenalty + parseLine(lines[i])
	}
	fmt.Println("Part 1: " + strconv.Itoa(totalPenalty))
}

func parseLine(line string) int {
	splitLine := strings.Split(line, "")
	done := false

	for !done {
		var right string
		var left string
		var rightIndex int

		for i := 0; i < len(splitLine); i++ {
			if strings.Contains(rightBraces, splitLine[i]) {
				right = splitLine[i]
				left = getMatchingLeft(right)
				rightIndex = i
				break
			}
		}

		if rightIndex <= 0 {
			return getPenalty(right)
		} else if splitLine[rightIndex-1] != left {
			return getPenalty(right)
		}

		splitLine = eraseElement(splitLine, rightIndex)
		splitLine = eraseElement(splitLine, rightIndex-1)
		if len(splitLine) <= 0 {
			done = true
		}
	}
	return 0
}

func getMatchingLeft(right string) string {
	if right == rightCurve {
		return leftCurve
	} else if right == rightSquare {
		return leftSquare
	} else if right == rightCurly {
		return leftCurly
	} else if right == rightDiamond {
		return leftDiamond
	}
	return "0"
}

func eraseElement(a []string, i int) []string {
	copy(a[i:], a[i+1:])
	a[len(a)-1] = ""
	return a[:len(a)-1]
}

func getPenalty(brace string) int {
	if brace == rightCurve {
		return 3
	} else if brace == rightSquare {
		return 57
	} else if brace == rightCurly {
		return 1197
	} else if brace == rightDiamond {
		return 25137
	}
	return 0
}

func readInput(path string) []string {
	f, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	var result []string

	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	return result
}
