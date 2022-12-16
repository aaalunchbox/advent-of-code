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
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	// Counters and Flags
	parseStacks := true

	// Slices representing stacks
	var crates [][]string

	// Read file, line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Err() != nil {
			log.Fatalln(scanner.Err())
		}
		var line = scanner.Text()

		// Top of the file are the stacks, parse stacks while in this part
		if parseStacks {
			// Check if blank line, end of stacks, start of moves
			if len(line) == 0 {
				fmt.Println("Empty line, stop parsing stacks")
				crates = transposeCrates(crates)
				parseStacks = false
				continue
			}

			if c := parseCrateLine(line); c != nil {
				crates = append(crates, c)
			}
			continue
		}

		// Done parsing stacks, continue to parsing moves
		move := parseMoveLine(line)
		crates = moveCrate(crates, move[0], move[1] - 1, move[2] - 1)
	}

	fmt.Println(crates)
	var answer string
	for _, v := range crates {
		answer += v[len(v) - 1]
	}
	fmt.Println("Answer:", answer)
}

func transposeCrates(crates [][]string) [][]string {
	t := make([][]string, len(crates[0]))
	for i := len(crates) - 1; i >= 0; i-- {
		for j := 0; j < len(crates[0]); j++ {
			if crates[i][j] == "" {
				continue
			}
			t[j] = append(t[j], crates[i][j])
		}
	}
	fmt.Println("Transpose:", t)
	return t
}

func moveCrate(crates [][]string, count int, from int, to int) [][]string {
	fmt.Println("Before:", crates)
	fmt.Println("Count:", count, "From:", from, "To", to)
	// If no crates in the from stack, nothing to do
	if len(crates[from]) == 0 {
		return crates
	}

	// If count is greater then length, set count to length since we can
	// only move available crates
	if count > len(crates[from]){
		fmt.Println("Count greater then available crates in FROM stack")
		count = len(crates[from])
	}

	// Pop crates from stack
	l := len(crates[from]) - count
	c := crates[from][l:]  // Get Crate
	crates[from] = crates[from][:l]  // Remove crate from stack

	// Append to TO stack
	crates[to] = append(crates[to], c...)
	fmt.Println("After:", crates)
	return crates
}

func parseCrateLine(line string) []string {
	var crate []string

	// Check if this is the last stack row, which is just column numbers
	l := []rune(strings.TrimSpace(line))
	if _, err := strconv.Atoi(string(l[0:1])); err == nil {
		fmt.Println("Column numbers, end of stacks")
		return nil
	}

	// Parse crate line
	for len(line) > 0 {
		var c string
		c, line = crateToken(line)
		crate = append(crate, c)
	}
	return crate
}

func crateToken(line string) (string, string) {
	var newLine string
	if len(line) > 3 {
		newLine = line[4:]
	} else {
		newLine = ""
	}
	fmt.Println("Token:", string(line[1]), "Remaining Line:", newLine)
	return strings.Trim(string(line[1]), " "), newLine
}

func parseMoveLine(line string) []int {
	// Split string by spaces
	s := strings.Split(line, " ")
	var x []int
	var m int
	m, _ = strconv.Atoi(s[1])
	x = append(x, m)
	m, _ = strconv.Atoi(s[3])
	x = append(x, m)
	m, _ = strconv.Atoi(s[5])
	x = append(x, m)
	return x
}
