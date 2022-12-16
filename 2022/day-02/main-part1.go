package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Define score constants
const win = 6
const lose = 0
const draw = 3

const rock = 1
const paper = 2
const scissors = 3

func main() {
	// Open file
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	// Score counter
	score := 0

	// Read file, line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Err() != nil {
			log.Fatalln(scanner.Err())
		}
		var line = scanner.Text()

		// Shouldn't be any blank lines, skip them just in case
		if len(line) == 0 {
			continue
		}

		v := strings.Fields(line)
		fmt.Println(v)
		score += calcScore(v)

		fmt.Print("\n\nFinal Score:", score, "\n")
	}
}

func calcScore(play []string) int {
	score := 0
	switch {
	case play[0] == "A":
		switch {
		case play[1] == "X":
			score = draw
		case play[1] == "Y":
			score = win
		case play[1] == "Z":
			score = lose
		default:
			log.Fatalln("Dunno what this is", play[1])
		}

	case play[0] == "B":
		switch {
		case play[1] == "X":
			score = lose
		case play[1] == "Y":
			score = draw
		case play[1] == "Z":
			score = win
		default:
			log.Fatalln("Dunno what this is", play[1])
		}

	case play[0] == "C":
		switch {
		case play[1] == "X":
			score = win
		case play[1] == "Y":
			score = lose
		case play[1] == "Z":
			score = draw
		default:
			log.Fatalln("Dunno what this is", play[1])
		}

	default:
		log.Fatalln("Dunno what this is", play[0])
	}

	fmt.Println("VS Score:", score)
	score += choiceScore(play[1])
	fmt.Println("Round Score:", score)
	return score
}

func choiceScore(choice string) int {
	score := 0
	switch {
	case choice == "X":
		score = rock
	case choice == "Y":
		score = paper
	case choice == "Z":
		score = scissors
	default:
		log.Fatalln("Dunno what this is", choice)
	}
	fmt.Println("Choice Score:", score)
	return score
}
