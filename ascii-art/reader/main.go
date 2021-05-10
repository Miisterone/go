package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"runtime"
	"strings"
)

var glines [8]string
var alphabet string

func main() {
	args := os.Args[1:]
	if len(args) != 1 || len(args[0]) < 1 { // Check if there is only one argument
		return
	}
	text := args[0]

	_, currentFilePath, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("Failed to find current dir")
	}
	currentDir := path.Dir(currentFilePath) // Get current dir path

	output, err := ioutil.ReadFile(path.Join(currentDir, "..", "fonts", "standard.txt")) // Read standard.txt file
	if err != nil {
		panic(err)
	}
	alphabet = string(output) // Stores the string conversion of the content in `alphabet`

	var carRtrns = make(map[int]bool, len(text)) // Makes a map that will store carriage return
	escaped := false
	for i := 0; i < len(text)-1; i++ {
		twoLength := text[i : i+2]
		if i < len(text)-2 {
			if twoLength == "\\\\" {
				text = text[:i] + text[i+1:]
				escaped = true // Protects double backslashes
				continue
			}
		}
		if twoLength == "\\n" {
			var bsn int
			if !escaped {
				j := i
				for j > 0 && text[j-1] == '\\' {
					bsn++ // Counts the number of backslashes
					j--
				}
			}
			escaped = false
			if bsn%2 == 0 {
				text = text[:i] + text[i+1:] // Removes a backslash if number of backslashes is even
				carRtrns[i] = true
			}
		}
	}

	for i, ltr := range text { // Iterates through given text
		if carRtrns[i] { // If there is a carriage return:
			printLines() // Displays the characters that have already been added to the lines
			glines = [8]string{} // Resets the lines
			continue
		}
		suffixLine(ltr) // Add current letter line by line to lines variables (line1, line2, ...)
		if i < len(text)-1 {
			for l := range glines { // Adds little separation between letters
				glines[l] += " "
			}
		}
	}
	printLines()
}

func getLtrPos(ltr rune) int { // Returns corresponding line number
	if ltr > 31 && ltr < 127 {
		return (int(ltr-32) * 9) + 2
	}
	return 0
}

func printLines() {
	for _, line := range glines { // Prints each line
		for _, c := range line { // Prints each char of the current line
			fmt.Print(string(c))
		}
		fmt.Print("\n")
	}
}

func readLine(lineNum int) string {
	r := strings.NewReader(alphabet) // Create new reader from file content
	sc := bufio.NewScanner(r)
	var lastLine int
	for sc.Scan() { // Looks for the good line
		lastLine++
		if lastLine == lineNum {
			return sc.Text()
		}
	}
	return "LINE NOT FOUND"
}

func suffixLine(ltr rune) {
	var lines [8]string
	lineNbr := getLtrPos(ltr) // Gets the number of the line corresponding to the char
	var maxLength int
	for i := range lines {
		lines[i] = readLine(lineNbr + i) // Gets each line of the letter and stoers it in the array
		if len(lines[i]) > maxLength {
			maxLength = len(lines[i]) // Gets the max length of a line of the char
		}
	}
	if ltr == ' ' { // Sets the max length to 5 if char is a space
		maxLength = 5
	}
	for i := range lines {
		for len(lines[i]) < maxLength { // Adds extra spaces if line length is inferior to max length
			lines[i] += " "
		}
	}
	for i := range glines {
		glines[i] += lines[i] // Adds letter lines to full text lines
	}
}
