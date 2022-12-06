package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Define score constants

func main() {
	// Open file
	file, err := os.Open("input-test.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	// Counters
	numOfOverlaps := 0

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

		// Save assignments to pair
		pair := createAssignment(line)
		fmt.Println(pair)

		// Check if one assignment is coverd by another
		if fullOverlap(pair) {
			fmt.Println("Overlap detected")
			numOfOverlaps += 1
		}
	}

	fmt.Println("Number of overlaps:", numOfOverlaps)
}

// Check if assignments fully overlap
func fullOverlap(a [][]int) bool {
	a1 := a[0]
	a2 := a[1]

	if a1[0] <= a2[0] && a1[1] >= a2[1] {
		return true
	}

	if a2[0] <= a1[0] && a2[1] >= a1[1] {
		return true
	}

	return false
}

// Convert a string to an assignment pair
func createAssignment(line string) [][]int {
	pair := [][]int{
		{0, 0},
		{0, 0},
	}
	pairSplit := strings.Split(line, ",")
	for i, v := range pairSplit {
		a := strings.Split(v, "-")
		a1, _ := strconv.Atoi(a[0])
		a2, _ := strconv.Atoi(a[1])
		pair[i][0] = a1
		pair[i][1] = a2
	}

	return pair
}
