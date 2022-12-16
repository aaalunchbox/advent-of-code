package main

import (
	"bufio"
	"container/list"
	"fmt"
	"log"
	"os"
	"unicode"
)

type Queue struct {
	*list.List
	Limit int
}

func (q *Queue) AddRune(r rune) {
	q.PushBack(r)

	// Check if limit is hit, throw out the first element
	if q.Len() > q.Limit {
		q.Remove(q.Front())
	}
}

func (q *Queue) AllUnique() bool {
	m := make(map[rune]rune)

	for e := q.Front(); e != nil; e = e.Next() {
		if _, ok := m[e.Value.(rune)]; ok {
			return false
		}
		m[e.Value.(rune)] = e.Value.(rune)
	}
	return true
}

func main() {
	// Open file
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	// Create a queue to store the last 4 runes
	q := Queue{
		list.New(),
		4,
	}

	// Read file, rune by rune
	reader := bufio.NewReader(file)
	loc := 0 // location of the current charcter
	for {
		r, _, err := reader.ReadRune()
		if err != nil {
			log.Fatal(err)
		}
		loc += 1

		// Skip empty runes
		if unicode.IsSpace(r) {
			fmt.Println("Empty rune")
			continue
		}

		// Add rune
		q.AddRune(r)

		// Keep going if queue isn't full
		if q.Len() < q.Limit {
			continue
		}

		if q.AllUnique() {
			fmt.Println("Found unique at", loc)
			break
		}
	}
}
