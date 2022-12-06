package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Define score constants

func main() {
	// Open file
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	// Counters
	totalPriority := 0
	lineCount := 0
	group := make([]string, 3)

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

		// Add line to group
		group[lineCount] = line

		// Check if we got 3 lines yet
		if lineCount < 2 {
			lineCount += 1
			continue
		}

		// Got 3 lines, reset count and process
		lineCount = 0
		fmt.Println(group)

		// Find duplicate and get priority
		badge := findBadge(group)
		fmt.Println("Duplicate:", string(badge))
		priority := calculatePriority(badge)
		fmt.Println("Priority:", priority)
		fmt.Print("\n\n")

		// Add priority to counter
		totalPriority += priority
	}

	fmt.Println("Total Priority:", totalPriority)
}

// Find the badge in all the bags
func findBadge(group []string) rune {
	for _, letter := range group[0] {
		if strings.Contains(group[1], string(letter)) && strings.Contains(group[2], string(letter)) {
			return letter
		}
	}
	return 0
}

// Convert a rune to a priority
func calculatePriority(item rune) int {
	if int(item) > 96 {
		return int(item) - 96
	} else {
		return int(item) - 38
	}

}
