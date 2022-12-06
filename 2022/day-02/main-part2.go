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
	}
	fmt.Print("\n\nFinal Score:", score, "\n")
}

func calcScore(play []string) int {
	score := 0
	switch {
	case play[1] == "X":
		// Lose
		score = lose
		move := choiceScore(play[0])
		if move == 1 {
			move = 3
		} else {
			move -= 1
		}
		fmt.Println("Move Score:", move)
		score += move

	case play[1] == "Y":
		// Draw
		score = draw
		move := choiceScore(play[0])
		fmt.Println("Move Score:", move)
		score += move

	case play[1] == "Z":
		// Win
		score = win
		move := choiceScore(play[0])
		if move == 3 {
			move = 1
		} else {
			move += 1
		}
		fmt.Println("Move Score:", move)
		score += move

	default:
		log.Fatalln("Dunno what scenario this is", play[1])
	}

	fmt.Println("Round Score:", score)
	return score
}

func choiceScore(choice string) int {
	score := 0
	switch {
	case choice == "A":
		score = rock
	case choice == "B":
		score = paper
	case choice == "C":
		score = scissors
	default:
		log.Fatalln("Dunno what choice this is", choice)
	}

	return score
}
