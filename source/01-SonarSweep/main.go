package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	part1()
	part2()
}

func part1() {

	f, err := os.Open("input-01.txt")

	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	var count int = 0

	for i := 0; i < (len(lines) - 1); i++ {
		value1, err1 := strconv.Atoi(lines[i])
		value2, err2 := strconv.Atoi(lines[i+1])
		if value2 > value1 {
			count = count + 1
		}

		if err1 != nil {
			fmt.Println(err1)
			log.Fatal(err1)
		}
		if err2 != nil {
			fmt.Println(err2)
			log.Fatal(err2)
		}
	}

	fmt.Println(count)
}

func part2() {
	f, err := os.Open("input-01.txt")

	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	var count int = 0
	var oldsum int = 0

	for i := 0; i < (len(lines) - 2); i++ {
		value1, err1 := strconv.Atoi(lines[i])
		value2, err2 := strconv.Atoi(lines[i+1])
		value3, err3 := strconv.Atoi(lines[i+2])

		var newsum int = value1 + value2 + value3

		if i > 0 {
			if newsum > oldsum {
				count = count + 1
			}

		}

		oldsum = newsum

		if err1 != nil {
			fmt.Println(err1)
			log.Fatal(err1)
		}
		if err2 != nil {
			fmt.Println(err2)
			log.Fatal(err2)
		}
		if err3 != nil {
			fmt.Println(err3)
			log.Fatal(err3)
		}
	}

	fmt.Println(count)
}
