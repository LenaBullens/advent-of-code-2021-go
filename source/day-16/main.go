package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var hexToBinMap = map[string]string{
	"0": "0000",
	"1": "0001",
	"2": "0010",
	"3": "0011",
	"4": "0100",
	"5": "0101",
	"6": "0110",
	"7": "0111",
	"8": "1000",
	"9": "1001",
	"A": "1010",
	"B": "1011",
	"C": "1100",
	"D": "1101",
	"E": "1110",
	"F": "1111",
}

func main() {
	part1()
}

func part1() {
	hex := readInput("input-16.txt")
	bin := hexToBin(hex)
	totalVersion := parsePacket(bin)
	fmt.Println("==========")
	fmt.Printf("Total version: %d\n", totalVersion)
}

func parsePacket(bin string) int {
	var totalVersion int

	cursor := 0

	//First 6 bits
	packetVersion, packetTypeId := parseHeader(bin[cursor : cursor+6])
	cursor = cursor + 6

	totalVersion = totalVersion + packetVersion

	if packetTypeId != 4 {
		//Parse length stuff
		lengthTypeID := bin[cursor : cursor+1]
		cursor = cursor + 1
		var length int
		if lengthTypeID == "0" {
			lengthBin := bin[cursor : cursor+15]
			length = binToInt(lengthBin)
			cursor = cursor + 15
			totalVersion = totalVersion + parseSubPacketsMode0(bin[cursor:cursor+length], length)
		} else {
			lengthBin := bin[cursor : cursor+11]
			length = binToInt(lengthBin)
			cursor = cursor + 11
			addedVersion, _ := parseSubPacketsMode1(bin[cursor:], length)
			totalVersion = totalVersion + addedVersion
		}
	}

	return totalVersion
}

func parseHeader(bin string) (int, int) {
	cursor := 0

	//First 3 bits
	packetVersionBin := bin[cursor : cursor+3]
	packetVersion := binToInt(packetVersionBin)
	cursor = cursor + 3
	fmt.Printf("Packet version: %d\n", packetVersion)

	//Next 3 bits
	packetTypeIDBin := bin[cursor : cursor+3]
	packetTypeID := binToInt(packetTypeIDBin)
	fmt.Printf("Packet ID: %d\n", packetTypeID)
	return packetVersion, packetTypeID
}

func parseSubPacketsMode0(bin string, totalLength int) int {
	var totalVersion int

	cursor := 0
	done := false

	for !done {
		//First 6 bits
		packetVersion, packetTypeId := parseHeader(bin[cursor : cursor+6])
		cursor = cursor + 6

		totalVersion = totalVersion + packetVersion

		if packetTypeId != 4 {
			//Parse length stuff
			lengthTypeID := bin[cursor : cursor+1]
			cursor = cursor + 1
			var length int
			if lengthTypeID == "0" {
				lengthBin := bin[cursor : cursor+15]
				length = binToInt(lengthBin)
				cursor = cursor + 15
				totalVersion = totalVersion + parseSubPacketsMode0(bin[cursor:cursor+length], length)
				cursor = cursor + length
			} else {
				lengthBin := bin[cursor : cursor+11]
				length = binToInt(lengthBin)
				cursor = cursor + 11
				addedVersion, cursorShift := parseSubPacketsMode1(bin[cursor:], length)
				totalVersion = totalVersion + addedVersion
				cursor = cursor + cursorShift
			}
		} else {
			doneLiteral := false
			for !doneLiteral {
				fiveBits := bin[cursor : cursor+5]
				oneBit := fiveBits[:1]
				cursor = cursor + 5
				if oneBit == "0" {
					doneLiteral = true
				}
			}
		}
		if cursor >= totalLength {
			done = true
		}

	}

	return totalVersion
}

func parseSubPacketsMode1(bin string, totalLength int) (int, int) {
	var totalVersion int

	cursor := 0
	nbPacketsParsed := 0
	done := false

	for !done {
		//First 6 bits
		packetVersion, packetTypeId := parseHeader(bin[cursor : cursor+6])
		cursor = cursor + 6
		nbPacketsParsed = nbPacketsParsed + 1

		totalVersion = totalVersion + packetVersion

		if packetTypeId != 4 {
			//Parse length stuff
			lengthTypeID := bin[cursor : cursor+1]
			cursor = cursor + 1
			var length int
			if lengthTypeID == "0" {
				lengthBin := bin[cursor : cursor+15]
				length = binToInt(lengthBin)
				cursor = cursor + 15
				totalVersion = totalVersion + parseSubPacketsMode0(bin[cursor:cursor+length], length)
				cursor = cursor + length
			} else {
				lengthBin := bin[cursor : cursor+11]
				length = binToInt(lengthBin)
				cursor = cursor + 11
				addedVersion, cursorShift := parseSubPacketsMode1(bin[cursor:], length)
				totalVersion = totalVersion + addedVersion
				cursor = cursor + cursorShift
			}
		} else {
			doneLiteral := false
			for !doneLiteral {
				fiveBits := bin[cursor : cursor+5]
				oneBit := fiveBits[:1]
				cursor = cursor + 5
				if oneBit == "0" {
					doneLiteral = true
				}
			}
		}
		if nbPacketsParsed >= totalLength {
			done = true
		}
	}

	return totalVersion, cursor
}

func hexToBin(hex string) string {
	var result string
	splitString := strings.Split(hex, "")
	for i := 0; i < len(splitString); i++ {
		result = result + hexToBinMap[splitString[i]]
	}
	return result
}

func binToInt(bin string) int {
	result, err := strconv.ParseInt(bin, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	return int(result)
}

func readInput(path string) string {
	f, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	var result string

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		result = result + scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}
