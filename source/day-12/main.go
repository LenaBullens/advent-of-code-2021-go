package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type cave struct {
	name string
	//What do if input has multiple links to same cave? Check in handling method???
	neighbours []string
	isStart    bool
	isEnd      bool
	isSmall    bool
}

var checkedPaths []string
var caveMap map[string]cave

func main() {
	part2()
}

func part1() {
	caveMap = readInput("input-12.txt")
	//Find the start cave.
	start := caveMap["start"]
	var path []cave
	step(start, path)
	for i := 0; i < len(checkedPaths); i++ {
		fmt.Println(checkedPaths[i])
	}
	fmt.Println(len(checkedPaths))
}

func part2() {
	caveMap = readInput("input-12.txt")
	//Find the start cave.
	start := caveMap["start"]
	var path []cave
	step2(start, path)
	//for i := 0; i < len(checkedPaths); i++ {
	//	fmt.Println(checkedPaths[i])
	//}
	fmt.Println(len(checkedPaths))
}

func step(cave cave, path []cave) {
	path = append(path, cave)
	//Check if end
	if cave.isEnd {
		pathString := pathToString(path)
		if !containsString(checkedPaths, pathString) {
			checkedPaths = append(checkedPaths, pathString)
		}
		return
	}

	//Check neighbours
	for i := 0; i < len(cave.neighbours); i++ {
		if caveMap[cave.neighbours[i]].isSmall {
			if !containsCave(caveSliceToStringSlice(path), cave.neighbours[i]) {
				step(caveMap[cave.neighbours[i]], path)
			}
		} else {
			step(caveMap[cave.neighbours[i]], path)
		}
	}
}

func step2(cave cave, path []cave) {
	path = append(path, cave)
	//Check if end
	if cave.isEnd {
		pathString := pathToString(path)
		if !containsString(checkedPaths, pathString) {
			fmt.Println(pathString)
			checkedPaths = append(checkedPaths, pathString)
		}
		return
	}

	//Check neighbours
	for i := 0; i < len(cave.neighbours); i++ {
		if caveMap[cave.neighbours[i]].isSmall {
			if canVisitSmallCave(caveSliceToStringSlice(path), cave.neighbours[i]) {
				step2(caveMap[cave.neighbours[i]], path)
			}
		} else {
			step2(caveMap[cave.neighbours[i]], path)
		}
	}
}

func caveSliceToStringSlice(caveSlice []cave) []string {
	var result []string
	for i := 0; i < len(caveSlice); i++ {
		result = append(result, caveSlice[i].name)
	}
	return result
}

func createCave(name string) cave {
	result := cave{}
	result.name = name
	if name == "start" {
		result.isStart = true
	}
	if name == "end" {
		result.isEnd = true
	}
	//Dirty trick to check if string starts with a lower-case letter.
	firstChar := name[0:1]
	if strings.ToLower(firstChar) == firstChar {
		result.isSmall = true
	}
	return result
}

func containsCave(caveSlice []string, cave string) bool {
	for _, c := range caveSlice {
		if c == cave {
			return true
		}
	}
	return false
}

func canVisitSmallCave(caveSlice []string, cave string) bool {
	if cave == "start" {
		return false
	}
	if !containsString(caveSlice, cave) {
		return true
	} else {
		if countOccurences(caveSlice, cave) >= 2 {
			return false
		} else {
			if checkIfAlreadyTwo(caveSlice, cave) {
				return false
			} else {
				return true
			}
		}
	}
}

func countOccurences(caveSlice []string, cave string) int {
	var count int
	for i := 0; i < len(caveSlice); i++ {
		if caveSlice[i] == cave {
			count = count + 1
		}
	}
	return count
}

func checkIfAlreadyTwo(caveSlice []string, cave string) bool {
	for i := 0; i < len(caveSlice); i++ {
		firstChar := caveSlice[i][0:1]
		if strings.ToLower(firstChar) == firstChar {
			var count int
			for j := 0; j < len(caveSlice); j++ {
				if caveSlice[i] == caveSlice[j] {
					count = count + 1
				}
			}
			if count >= 2 {
				return true
			}
		}
	}
	return false
}

func containsString(stringSlice []string, str string) bool {
	for _, s := range stringSlice {
		if s == str {
			return true
		}
	}
	return false
}

func pathToString(path []cave) string {
	var result string
	for i := 0; i < len(path); i++ {
		result = result + path[i].name
		if i != len(path)-1 {
			result = result + ","
		}
	}
	return result
}

func readInput(path string) map[string]cave {
	f, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	caveMap := make(map[string]cave)

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		splitLine := strings.Split(scanner.Text(), "-")
		caveName1 := splitLine[0]
		caveName2 := splitLine[1]
		var cave1 cave
		var cave2 cave
		//Check if first cave already in map
		_, present1 := caveMap[caveName1]
		if !present1 {
			cave1 = createCave(caveName1)
		} else {
			cave1 = caveMap[caveName1]
		}
		//Check if second cave already in map
		_, present2 := caveMap[caveName2]
		if !present2 {
			cave2 = createCave(caveName2)
		} else {
			cave2 = caveMap[caveName2]
		}
		//Setup cave links if not present yet
		if !containsCave(cave1.neighbours, caveName2) {
			cave1.neighbours = append(cave1.neighbours, caveName2)
		}
		if !containsCave(cave2.neighbours, caveName1) {
			cave2.neighbours = append(cave2.neighbours, caveName1)
		}

		caveMap[caveName1] = cave1
		caveMap[caveName2] = cave2
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return caveMap
}
