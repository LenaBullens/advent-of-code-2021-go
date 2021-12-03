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
}
