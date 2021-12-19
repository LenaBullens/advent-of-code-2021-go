package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	part1()
	part2()
}

func part1() {

	f, err := os.Open("input-03.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	var lines [][]int

	for scanner.Scan() {
		split := strings.Split(scanner.Text(), "")
		var line []int
		for i := 0; i < (len(split)); i++ {
			bit, err := strconv.Atoi(split[i])
			if err != nil {
				log.Fatal(err)
			}
			line = append(line, bit)
		}
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	width := len(lines[0])
	var oneBits = make([]int, width)

	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			if lines[i][j] == 1 {
				oneBits[j] = oneBits[j] + 1
			}
		}
	}

	var binGamma string
	var binEpsilon string

	for i := 0; i < len(oneBits); i++ {
		if oneBits[i] > (len(lines) - oneBits[i]) {
			binGamma = binGamma + "1"
			binEpsilon = binEpsilon + "0"
		} else {
			binGamma = binGamma + "0"
			binEpsilon = binEpsilon + "1"
		}
	}

	gamma, err := strconv.ParseInt(binGamma, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	epsilon, err := strconv.ParseInt(binEpsilon, 2, 64)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(gamma * epsilon)
}

func part2() {

	f, err := os.Open("input-03.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	var lines [][]int

	for scanner.Scan() {
		split := strings.Split(scanner.Text(), "")
		var line []int
		for i := 0; i < (len(split)); i++ {
			bit, err := strconv.Atoi(split[i])
			if err != nil {
				log.Fatal(err)
			}
			line = append(line, bit)
		}
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var oxygenLines = make([][]int, len(lines))
	copy(oxygenLines, lines)
	var oxygenString = calculateOxygen(oxygenLines)

	var co2Lines = make([][]int, len(lines))
	copy(co2Lines, lines)
	var co2String = calculateCO2(co2Lines)

	oxygen, err := strconv.ParseInt(oxygenString, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	co2, err := strconv.ParseInt(co2String, 2, 64)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(oxygen * co2)
}

func calculateOxygen(lines [][]int) string {
	var width int = len(lines[0])
	for i := 0; i < width; i++ {
		var zeroes int = 0
		var ones int = 0
		for j := 0; j < len(lines); j++ {
			if lines[j][i] == 1 {
				ones = ones + 1
			} else {
				zeroes = zeroes + 1
			}
		}

		if ones >= zeroes {
			var newSlice = make([][]int, 0)
			for j := 0; j < len(lines); j++ {
				if lines[j][i] == 1 {
					newSlice = append(newSlice, lines[j])
				}
			}
			lines = newSlice
		} else {
			var newSlice = make([][]int, 0)
			for j := 0; j < len(lines); j++ {
				if lines[j][i] == 0 {
					newSlice = append(newSlice, lines[j])
				}
			}
			lines = newSlice
		}

		if len(lines) == 1 {
			break
		}
	}

	var oxygen string
	for i := 0; i < len(lines[0]); i++ {
		oxygen = oxygen + strconv.Itoa(lines[0][i])
	}
	return oxygen
}

func calculateCO2(lines [][]int) string {
	var width int = len(lines[0])
	for i := 0; i < width; i++ {
		var zeroes int = 0
		var ones int = 0
		for j := 0; j < len(lines); j++ {
			if lines[j][i] == 1 {
				ones = ones + 1
			} else {
				zeroes = zeroes + 1
			}
		}

		if ones >= zeroes {
			var newSlice = make([][]int, 0)
			for j := 0; j < len(lines); j++ {
				if lines[j][i] == 0 {
					newSlice = append(newSlice, lines[j])
				}
			}
			lines = newSlice
		} else {
			var newSlice = make([][]int, 0)
			for j := 0; j < len(lines); j++ {
				if lines[j][i] == 1 {
					newSlice = append(newSlice, lines[j])
				}
			}
			lines = newSlice
		}

		if len(lines) == 1 {
			break
		}
	}

	var co2 string
	for i := 0; i < len(lines[0]); i++ {
		co2 = co2 + strconv.Itoa(lines[0][i])
	}
	return co2
}
