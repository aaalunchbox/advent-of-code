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

	// Priority Counter
	totalPriority := 0

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
		fmt.Println(line)
		duplicate := findDuplicate(splitSack(line))
		fmt.Println("Duplicate:", string(duplicate))
		priority := calculatePriority(duplicate)
		fmt.Println("Priority:", priority)
		fmt.Print("\n\n")

		// Add priority to counter
		totalPriority += priority
	}

	fmt.Println("Total Priority:", totalPriority)
}

// Split the content line of a sack into two even strings
func splitSack(line string) (string, string) {
	half := len(line) / 2
	c1 := line[0:half]
	c2 := line[half:]
	if len(c1) != len(c2) {
		log.Fatalln("Content split not equal: ", line)
	}
	return c1, c2
}

// Find dupliacte items between the bags
func findDuplicate(c1 string, c2 string) rune {
	for _, letter := range c1 {
		if strings.Contains(c2, string(letter)) {
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
