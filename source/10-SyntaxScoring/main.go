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
	part2()
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

func part2() {
	lines := readInput("input-10.txt")
	//The only missing parantheses are right ones. We look for the right most left one. The one to the
	//right of it should be its matching parenthesis. If we can't find one, it has to be added. Remove
	//and repeat.

	var scores []int
	var uncorruptedLines []string

	for i := 0; i < len(lines); i++ {
		if parseLine(lines[i]) <= 0 {
			uncorruptedLines = append(uncorruptedLines, lines[i])
		}
	}

	for i := 0; i < len(uncorruptedLines); i++ {
		result := completeLine(uncorruptedLines[i])
		score := additionToScore(result)
		scores = append(scores, score)
	}

	sort.Ints(scores)

	fmt.Println("Part 2: " + strconv.Itoa(scores[(len(scores)-1)/2]))
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

func completeLine(line string) string {
	splitLine := strings.Split(line, "")
	done := false
	var result string

	for !done {
		var right string
		var left string
		var leftIndex int

		for i := len(splitLine) - 1; i >= 0; i-- {
			if strings.Contains(leftBraces, splitLine[i]) {
				left = splitLine[i]
				right = getMatchingRight(left)
				leftIndex = i
				break
			}
		}

		//It's the right-most character => Add match to result.
		if leftIndex >= len(splitLine)-1 {
			result = result + right
			splitLine = eraseElement(splitLine, leftIndex)
		} else {
			splitLine = eraseElement(splitLine, leftIndex)
			//Right paranthese shifted one to the left because of erase, so no index + 1.
			splitLine = eraseElement(splitLine, leftIndex)
		}
		if len(splitLine) <= 0 {
			done = true
		}
	}
	return result
}

func additionToScore(addition string) int {
	splitAddition := strings.Split(addition, "")
	var score int
	for i := 0; i < len(splitAddition); i++ {
		score = score * 5
		score = score + getScore(splitAddition[i])
	}
	return score
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

func getMatchingRight(left string) string {
	if left == leftCurve {
		return rightCurve
	} else if left == leftSquare {
		return rightSquare
	} else if left == leftCurly {
		return rightCurly
	} else if left == leftDiamond {
		return rightDiamond
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

func getScore(brace string) int {
	if brace == rightCurve {
		return 1
	} else if brace == rightSquare {
		return 2
	} else if brace == rightCurly {
		return 3
	} else if brace == rightDiamond {
		return 4
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
