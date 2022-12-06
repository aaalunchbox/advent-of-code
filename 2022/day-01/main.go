package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	// Open file
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	// Create a variable to store and track elf calories
	topCalories := make([]int, 3)
	totalCalories := 0

	// Read file, line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Err() != nil {
			log.Fatalln(scanner.Err())
		}
		var line = scanner.Text()

		// If line is blank, new elf, check for new max and reset calorie counter
		if len(line) == 0 {
			updateTopThree(topCalories, totalCalories)
			totalCalories = 0
			continue
		}

		// Sum calories
		calories, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalln(err)
		}
		totalCalories += calories
	}

	// Check one more time if total calories is a new max, to get over EOF errors
	updateTopThree(topCalories, totalCalories)

	fmt.Println("Top 3 elf calories:", topCalories)

	calTotals := 0
	for _, v := range topCalories {
		calTotals += v
	}

	fmt.Println("Sum of calories per elf:", calTotals)
}

func updateTopThree(totals []int, calories int) {
	// Find the smallest calories in the totals
	minIdx := 0
	min := totals[minIdx]
	for i, v := range totals {
		if v < min {
			min = v
			minIdx = i
		}
	}
	// Check if calories is bigger, replace if it is
	if calories > min {
		totals[minIdx] = calories
	}
}
