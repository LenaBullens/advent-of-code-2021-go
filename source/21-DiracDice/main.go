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

type gamestate struct {
	pos0   int
	pos1   int
	score0 int
	score1 int
}

func main() {
	solve1()
	solve2()
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

/**
There are too many universes to simulate all of them. We can make several observations though.
	1. The history of a universe doesn't it's further progression. This means that two universes in
	which both players have the same score, position, and the same player is next are essentially
	interchangeable. In other words, there are 10 * 10 * 21 * 21 * 2 distinguishable states a universe
	can be in without a player having won.const
	2. 3d3 has 27 different outcomes, but the total value only ranges from 3 to 9, with varying frequencies
	for each result.
		3 : 1
		4 : 111
		5 : 111111
		6 : 1111111
		7 : 111111
		8 : 111
		9 : 1
So we can just keep track of how many universes exist in each state, and each turn create the required new universes
and remove the old ones. We meet again, Mister Laternfish.
**/

func solve2() {
	freqTable := map[int]int{
		3: 1,
		4: 3,
		5: 6,
		6: 7,
		7: 6,
		8: 3,
		9: 1,
	}
	pos := readInput("input-21.txt")
	currentPlayer := 0
	gamestates := make(map[gamestate]int)
	initialState := gamestate{pos0: pos[0], pos1: pos[1], score0: 0, score1: 0}
	gamestates[initialState] = 1
	done := false

	wins0 := 0
	wins1 := 0

	for !done {
		newGamestates := make(map[gamestate]int)
		for g, f := range gamestates {
			if f > 0 {
				//Create new gamestates
				for v, vf := range freqTable {
					posShift := v
					if currentPlayer == 0 {
						newPos := g.pos0 + posShift
						//New position is value smaller than 20. 'Furthest' jump is 10 (pos) + 9 (die).
						if newPos > 10 {
							newPos = newPos - 10
						}
						newScore := g.score0 + newPos
						newFreq := f * vf
						if newScore >= 21 {
							//Winner!
							wins0 = wins0 + newFreq
						} else {
							newg := gamestate{pos0: newPos, pos1: g.pos1, score0: newScore, score1: g.score1}
							newGamestates[newg] = newGamestates[newg] + newFreq
						}
					} else {
						newPos := g.pos1 + posShift
						if newPos > 10 {
							newPos = newPos - 10
						}
						newScore := g.score1 + newPos
						newFreq := f * vf
						if newScore >= 21 {
							wins1 = wins1 + newFreq
						} else {
							newg := gamestate{pos0: g.pos0, pos1: newPos, score0: g.score0, score1: newScore}
							newGamestates[newg] = newGamestates[newg] + newFreq
						}
					}
				}
			}
		}
		gamestates = newGamestates
		currentPlayer = 1 - currentPlayer
		//Check if any non-winning states are left.
		statesLeft := false
		for _, f := range gamestates {
			if f > 0 {
				statesLeft = true
				break
			}
		}
		done = !statesLeft
	}
	fmt.Printf("Player 1 wins: %d\n", wins0)
	fmt.Printf("Player 2 wins: %d\n", wins1)
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
