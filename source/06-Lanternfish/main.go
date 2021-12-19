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
	fmt.Println(part1())
	fmt.Println(part2())
}

func part1() int {
	fish := readInput("input-06.txt")
	for i := 0; i < 80; i++ {
		var fishToAdd int = 0
		for j := 0; j < len(fish); j++ {
			if fish[j] > 0 {
				fish[j] = fish[j] - 1
			} else if fish[j] == 0 {
				fishToAdd++
				fish[j] = 6
			}
		}
		for j := 0; j < fishToAdd; j++ {
			fish = append(fish, 8)
		}
	}

	return len(fish)
}

func part2() int {
	fish := readInput("input-06.txt")
	//Fish are indistinguishable outside of age. Any two fish of the same age are interchangeable. Thus, we can
	//just consider the population distribution instead of tracking each individual fish.
	var nbOfFishByAge [9]int
	for i := 0; i < len(fish); i++ {
		nbOfFishByAge[fish[i]] = nbOfFishByAge[fish[i]] + 1
	}

	for i := 0; i < 256; i++ {
		//We just need to shift fish between age 'columns' and add new fish. When fish leave column 0 they
		//get added to both column 6 and column 8. (Newborn fish + reproduction cooldown)
		var newNbOfFishByAge [9]int

		newNbOfFishByAge[0] = nbOfFishByAge[1]
		newNbOfFishByAge[1] = nbOfFishByAge[2]
		newNbOfFishByAge[2] = nbOfFishByAge[3]
		newNbOfFishByAge[3] = nbOfFishByAge[4]
		newNbOfFishByAge[4] = nbOfFishByAge[5]
		newNbOfFishByAge[5] = nbOfFishByAge[6]
		newNbOfFishByAge[6] = nbOfFishByAge[7] + nbOfFishByAge[0]
		newNbOfFishByAge[7] = nbOfFishByAge[8]
		newNbOfFishByAge[8] = nbOfFishByAge[0]
		nbOfFishByAge = newNbOfFishByAge
	}

	var totalFish int
	for i := 0; i < 9; i++ {
		totalFish = totalFish + nbOfFishByAge[i]
	}

	return totalFish
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

	fishAsStrings := strings.Split(input, ",")
	var fish []int

	for i := 0; i < len(fishAsStrings); i++ {
		fishValue, err := strconv.Atoi(fishAsStrings[i])
		if err != nil {
			log.Fatal(err)
		}
		fish = append(fish, fishValue)
	}
	return fish
}
