package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type deterministicDie struct {
	value       int
	timesRolled int
}

func main() {
	solve1()
}

func solve1() {
	pos := readInput("input-21.txt")
	currentPlayer := 0
	die := createDeterministicDie()
	donePlaying := false
	score := [2]int{0, 0}
	winner := 0

	for !donePlaying {
		for i := 0; i < 3; i++ {
			die = rollDie(die)
			pos[currentPlayer] = (pos[currentPlayer] + die.value) % 10
		}
		if pos[currentPlayer] == 0 {
			score[currentPlayer] = score[currentPlayer] + 10
		} else {
			score[currentPlayer] = score[currentPlayer] + pos[currentPlayer]

		}
		if score[currentPlayer] >= 1000 {
			donePlaying = true
			winner = currentPlayer
		}
		currentPlayer = 1 - currentPlayer
	}
	result := score[1-winner] * die.timesRolled
	fmt.Printf("Result: %d\n", result)
}

func createDeterministicDie() deterministicDie {
	return deterministicDie{value: 100, timesRolled: 0}
}

func rollDie(die deterministicDie) deterministicDie {
	die.timesRolled = die.timesRolled + 1
	if die.value >= 100 {
		die.value = 1
	} else {
		die.value = die.value + 1
	}
	return die
}

func readInput(path string) []int {
	f, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var pos []int

	for scanner.Scan() {
		splitLine := strings.Split(scanner.Text(), ": ")
		value, err := strconv.Atoi(splitLine[1])
		if err != nil {
			log.Fatal(err)
		}
		pos = append(pos, value)

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return pos
}
