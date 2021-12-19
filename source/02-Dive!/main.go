package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type step struct {
	direction string
	units     int
}

func main() {
	part1()
	part2()
}

func part1() {

	f, err := os.Open("input-02.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	var lines []step

	for scanner.Scan() {
		split := strings.Split(scanner.Text(), " ")
		dir := split[0]
		dis, err := strconv.Atoi(split[1])
		if err != nil {
			log.Fatal(err)
		}
		s := step{dir, dis}
		lines = append(lines, s)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var horPos int = 0
	var verPos int = 0

	for i := 0; i < (len(lines)); i++ {
		if lines[i].direction == "forward" {
			horPos = horPos + lines[i].units
		} else if lines[i].direction == "up" {
			verPos = verPos - lines[i].units
		} else if lines[i].direction == "down" {
			verPos = verPos + lines[i].units
		}
	}

	fmt.Println(horPos * verPos)
}

func part2() {

	f, err := os.Open("input-02.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	var lines []step

	for scanner.Scan() {
		split := strings.Split(scanner.Text(), " ")
		dir := split[0]
		dis, err := strconv.Atoi(split[1])
		if err != nil {
			log.Fatal(err)
		}
		s := step{dir, dis}
		lines = append(lines, s)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var horPos int = 0
	var verPos int = 0
	var aim int = 0

	for i := 0; i < (len(lines)); i++ {
		if lines[i].direction == "forward" {
			horPos = horPos + lines[i].units
			verPos = verPos + aim*lines[i].units
		} else if lines[i].direction == "up" {
			aim = aim - lines[i].units
		} else if lines[i].direction == "down" {
			aim = aim + lines[i].units
		}
	}

	fmt.Println(horPos * verPos)
}
