package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Directory struct {
	Name   string
	Parent *Directory
	Files  []File
	Dirs   []*Directory
}

type File struct {
	Name string
	Size int
}

func NewDirectory(name string, parent *Directory) *Directory{
	return &Directory{
		Name:   name,
		Parent: parent,
	}
} 

func main() {
	// Open file
	file, err := os.Open("input-test.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	// Create directory, starting with root
	dirs := NewDirectory("/", nil)

	cd := &dirs
	lastCmd := ""

	// Read file, line by line and build directory
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Err() != nil {
			log.Fatalln(scanner.Err())
		}
		line := scanner.Text()
		tokens := strings.Fields(line)

		if len(tokens) == 0 {
			continue
		}

		if tokens[0] == "$" {
			lastCmd = UserCommand(tokens, dirs, cd)
		} else if lastCmd == "ls" {
			AddDirContent(tokens, dirs)
		} else {
			log.Fatal("Unknown input")
		}

		fmt.Println(line)
	}
}

func UserCommand(tokens []string, dirs Directory, cd *Directory) string {
	fmt.Println("User Command:", tokens[1])
	return tokens[1]
}

func AddDirContent(tokens []string, dirs Directory) {

}
