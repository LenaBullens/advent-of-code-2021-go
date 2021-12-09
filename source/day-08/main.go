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

type display struct {
	signalPatterns []string
	outputValues   []string
}

func main() {
	fmt.Println(part1(readInput("input-08.txt")))
	fmt.Println(part2(readInput("input-08.txt")))
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

func part2(displays []display) int {
	var totalScore int
	for i := 0; i < len(displays); i++ {
		totalScore = totalScore + solveLine(displays[i])
	}
	return totalScore
}

func solveLine(display display) int {
	var nbOfLetterOccurences [7]int
	for i := 0; i < 10; i++ {
		if strings.Contains(display.signalPatterns[i], "a") {
			nbOfLetterOccurences[0] = nbOfLetterOccurences[0] + 1
		}
		if strings.Contains(display.signalPatterns[i], "b") {
			nbOfLetterOccurences[1] = nbOfLetterOccurences[1] + 1
		}
		if strings.Contains(display.signalPatterns[i], "c") {
			nbOfLetterOccurences[2] = nbOfLetterOccurences[2] + 1
		}
		if strings.Contains(display.signalPatterns[i], "d") {
			nbOfLetterOccurences[3] = nbOfLetterOccurences[3] + 1
		}
		if strings.Contains(display.signalPatterns[i], "e") {
			nbOfLetterOccurences[4] = nbOfLetterOccurences[4] + 1
		}
		if strings.Contains(display.signalPatterns[i], "f") {
			nbOfLetterOccurences[5] = nbOfLetterOccurences[5] + 1
		}
		if strings.Contains(display.signalPatterns[i], "g") {
			nbOfLetterOccurences[6] = nbOfLetterOccurences[6] + 1
		}
	}

	//Bottom-left segment occurs 4 times (uniquely).
	blSegment := getLetterForNumber(getIndexOfOccurence(nbOfLetterOccurences, 4)[0])

	//Top-left segment occurs 6 times (uniquely).
	tlSegment := getLetterForNumber(getIndexOfOccurence(nbOfLetterOccurences, 6)[0])

	//Bottom-right segment occurs 9 times (uniquely).
	brSegment := getLetterForNumber(getIndexOfOccurence(nbOfLetterOccurences, 9)[0])

	//To identify top segment we can look at 1 (2 seg) and 7 (3 seg)
	oneString := getStringsOfLength(display, 2)[0]
	oneStringChars := strings.Split(oneString, "")
	sevenString := getStringsOfLength(display, 3)[0]
	for i := 0; i < 2; i++ {
		sevenString = strings.ReplaceAll(sevenString, oneStringChars[i], "")
	}
	tSegment := sevenString

	//1 is built up of the top-right and top-left segments. We know which string represents 1
	//and know the mapping for one segment, se we can figure out the last one.
	trSegment := strings.ReplaceAll(oneString, brSegment, "")

	//We know the bottom-left and top-right segments. The only digit that contains 6 segments that
	//contains both is 0.
	sixSegmentStrings := getStringsOfLength(display, 6)
	var zeroString string
	for i := 0; i < 3; i++ {
		if strings.Contains(sixSegmentStrings[i], blSegment) && strings.Contains(sixSegmentStrings[i], trSegment) {
			zeroString = sixSegmentStrings[i]
		}
	}
	//The segment not in the zeroString is the middle segment.
	allSegments := "abcdefg"
	zeroStringChars := strings.Split(zeroString, "")
	for i := 0; i < 6; i++ {
		allSegments = strings.ReplaceAll(allSegments, zeroStringChars[i], "")
	}
	mSegment := allSegments

	//Bottom segment is the last remaining segment.
	allSegments = "abcdefg"
	allSegments = strings.ReplaceAll(allSegments, tSegment, "")
	allSegments = strings.ReplaceAll(allSegments, tlSegment, "")
	allSegments = strings.ReplaceAll(allSegments, trSegment, "")
	allSegments = strings.ReplaceAll(allSegments, mSegment, "")
	allSegments = strings.ReplaceAll(allSegments, blSegment, "")
	allSegments = strings.ReplaceAll(allSegments, brSegment, "")
	bSegment := allSegments

	var value string
	for i := 0; i < 4; i++ {
		var translatedString string
		outputValueChars := strings.Split(display.outputValues[i], "")
		for j := 0; j < len(outputValueChars); j++ {
			if outputValueChars[j] == tSegment {
				translatedString = translatedString + "a"
			} else if outputValueChars[j] == tlSegment {
				translatedString = translatedString + "b"
			} else if outputValueChars[j] == trSegment {
				translatedString = translatedString + "c"
			} else if outputValueChars[j] == mSegment {
				translatedString = translatedString + "d"
			} else if outputValueChars[j] == blSegment {
				translatedString = translatedString + "e"
			} else if outputValueChars[j] == brSegment {
				translatedString = translatedString + "f"
			} else if outputValueChars[j] == bSegment {
				translatedString = translatedString + "g"
			}
		}
		value = value + stringToDigit(sortLettersInStringAlphabetically(translatedString))
	}
	result, err := strconv.Atoi(value)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func stringToDigit(input string) string {
	if input == "abcefg" {
		return "0"
	} else if input == "cf" {
		return "1"
	} else if input == "acdeg" {
		return "2"
	} else if input == "acdfg" {
		return "3"
	} else if input == "bcdf" {
		return "4"
	} else if input == "abdfg" {
		return "5"
	} else if input == "abdefg" {
		return "6"
	} else if input == "acf" {
		return "7"
	} else if input == "abcdefg" {
		return "8"
	} else {
		return "9"
	}
}

func getIndexOfOccurence(array [7]int, occ int) []int {
	var indexes []int
	for i := 0; i < 7; i++ {
		if occ == array[i] {
			indexes = append(indexes, i)
		}
	}
	return indexes
}

func getStringsOfLength(display display, length int) []string {
	var output []string
	for i := 0; i < 10; i++ {
		if len(display.signalPatterns[i]) == length {
			output = append(output, display.signalPatterns[i])
		}
	}
	return output
}

func getLetterForNumber(nb int) string {
	if nb == 0 {
		return "a"
	} else if nb == 1 {
		return "b"
	} else if nb == 2 {
		return "c"
	} else if nb == 3 {
		return "d"
	} else if nb == 4 {
		return "e"
	} else if nb == 5 {
		return "f"
	} else if nb == 6 {
		return "g"
	}
	return ""
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
		//Sort the letters in the strings alphabetically to make things easier later on.
		for i := 0; i < 10; i++ {
			display.signalPatterns[i] = sortLettersInStringAlphabetically(display.signalPatterns[i])
		}
		for i := 0; i < 4; i++ {
			display.outputValues[i] = sortLettersInStringAlphabetically(display.outputValues[i])
		}
		displays = append(displays, display)
	}

	return displays
}

func sortLettersInStringAlphabetically(input string) string {
	letters := strings.Split(input, "")
	sort.Strings(letters)
	var output string
	for i := 0; i < len(letters); i++ {
		output = output + letters[i]
	}
	return output
}
