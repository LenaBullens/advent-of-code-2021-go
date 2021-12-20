package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	part1()
	part2()
}

func part1() {
	polymer, rules := readInput("input-14.txt")
	for i := 0; i < 40; i++ {
		fmt.Print("Step ")
		fmt.Println(i + 1)
		polymer = step1(polymer, rules)
	}

	occurences := make(map[string]int)

	for i := 0; i < len(polymer); i++ {
		occurences[polymer[i:i+1]] = occurences[polymer[i:i+1]] + 1
	}

	var least int = len(polymer)
	var most int = -1

	for _, value := range occurences {
		if value <= least {
			least = value
		}
		if value >= most {
			most = value
		}
	}
	difference := most - least
	fmt.Println(difference)
}

func part2() {
	polymer, rules := readInput("input-14.txt")
	pairs := make(map[string]int)

	for i := 0; i < len(polymer)-1; i++ {
		segment := polymer[i : i+2]
		pairs[segment] = pairs[segment] + 1
	}

	for i := 0; i < 40; i++ {
		newPairs := make(map[string]int)
		for pair, value := range pairs {
			var newPair1 string
			var newPair2 string
			newPair1 = pair[:1] + rules[pair]
			newPair2 = rules[pair] + pair[1:]
			newPairs[newPair1] = newPairs[newPair1] + value
			newPairs[newPair2] = newPairs[newPair2] + value
		}
		pairs = newPairs
	}

	occurences := make(map[string]int)

	//first and last symbol in starting polymer are only counted once in pairs, while all other ones
	//are counted twice, so we add them once.
	occurences[polymer[:1]] = 1
	occurences[polymer[len(polymer)-1:]] = 1

	for pair, value := range pairs {
		firstSymbol := pair[:1]
		secondSymbol := pair[1:]
		occurences[firstSymbol] = occurences[firstSymbol] + value
		occurences[secondSymbol] = occurences[secondSymbol] + value
	}

	var least int
	var most int
	isFirst := true

	for _, value := range occurences {
		if isFirst {
			least = value
			most = value
			isFirst = false
		}
		if value <= least {
			least = value
		}
		if value >= most {
			most = value
		}
	}
	//Divide by 2 cause we counted each symbol in two pairs.
	difference := (most - least) / 2
	fmt.Println(difference)
}

func step1(polymer string, rules map[string]string) string {
	var result string
	for i := 0; i < len(polymer)-1; i++ {
		segment := polymer[i : i+2]
		insert, exists := rules[segment]
		if exists {
			result = result + segment[:1] + insert
		} else {
			result = result + segment[:1]
		}
	}
	result = result + polymer[len(polymer)-1:]
	return result
}

func readInput(path string) (string, map[string]string) {
	mode := 0
	f, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	var polymer string
	rules := make(map[string]string)

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		if mode == 0 {
			if line != "" {
				polymer = line
			} else {
				mode = 1
			}
		} else {
			linePart1 := line[:2]
			linePart2 := line[6:]
			rules[linePart1] = linePart2
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return polymer, rules
}
