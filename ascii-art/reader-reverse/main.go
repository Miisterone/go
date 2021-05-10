package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"reflect"
	"runtime"
	"strconv"
	"strings"
)

var file string
var alphabet string
var alphaLength int
var doDebug bool

func main() {
	var err error
	var input []byte
	var output []byte
	var result string

	args := os.Args[1:]
	if len(args) > 3 || len(args) < 1 { // Check if there is only one argument
		return
	}

	var debugStr string
	SetStringFlag(&debugStr, "debug", "false", &args)
	doDebug, err = strconv.ParseBool(debugStr)
	if err != nil {
		doDebug = false
	}

	var reverse string
	SetStringFlag(&reverse, "reverse", "", &args) // Parse |reverse` flag
	if reverse == "" {
		return
	}

	if len(args) == 0 {
		args = append(args, "standard") // Sets standard font if no font is provided
	}

	_, currentFilePath, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("Failed to find current dir")
	}
	currentDir := path.Dir(currentFilePath) // Gets current dir path

	input, err = ioutil.ReadFile(path.Join(currentDir, "inputs", reverse))
	if err != nil {
		log.Fatal(errors.New("file'" + reverse + "'not found"))
	}
	file = string(input)

	if args[0] == "shadow" {
		output, err = ioutil.ReadFile(path.Join(currentDir, "..", "fonts", "shadow.txt")) // Reads shadow.txt file
	} else if args[0] == "thinkertoy" {
		output, err = ioutil.ReadFile(path.Join(currentDir, "..", "fonts", "thinkertoy.txt")) // Reads thinkertoy.txt file
	} else {
		output, err = ioutil.ReadFile(path.Join(currentDir, "..", "fonts", "standard.txt")) // Reads standard.txt file
	}
	if err != nil {
		panic(err)
	}
	alphabet = string(output)
	alphaLength, err = countLines(strings.NewReader(alphabet)) // Gets the number of lines in the font file
	if err != nil {
		panic(err)
	}

	linesNbr, err := countLines(strings.NewReader(file)) // Gets the number of lines in the input file
	if err != nil {
		panic(err)
	}
	paraNbr := linesNbr / 8 // Gets the number of "Ascii-Art lines" (a line of 8 high)

	var shortLetters = getShortLetters(3) // Gets every letter with a width of 3 or less

	for p := 0; p < paraNbr; p++ {
		if p > 0 {
			result += "\n" // Adds a carriage return if a line is ended
		}
		var potentialsLtrs []rune
		maxLineLength := getMaxLength(&file, p*8) // Gets the maximal length of the Ascii-Art line
		if maxLineLength < 0 {
			log.Fatal("Tried to get length of undefined line")
		}
		var ltrCol, totalLength int
		for c := 0; c < maxLineLength; c++ {
			if isSpaceFrom(file, c, p*8) { // Check if there is a space
				if len(potentialsLtrs) == 0 { // Must be no
					result += " "
					totalLength += 6
					c += 5
					continue
				} else {
					log.Fatal("Unexpected space at position " + strconv.Itoa(c))
				}
			}
			col, err := readCol(file, c, p*8) // Gets column from input
			if err != nil {
				log.Fatal("Failed to read col", col, "below line", p*8)
			}
			if isColEmpty(col) {
				var candidats = getIntersec(shortLetters, potentialsLtrs) // Gets intersection between small-width letters and potential letters
				if len(candidats) == 1 {
					potentialsLtrs = []rune{candidats[0]}
				} else {
					totalLength++
					continue
				}
			} else {
				getPossibilities(ltrCol, col, &potentialsLtrs) // Gets every possible rune that match
			}
			if len(potentialsLtrs) == 1 { // It's a match !
				// Let's go baby !
				result += string(potentialsLtrs[0])
				debug("Pre-result:", result)
				totalLength += getMaxLength(&alphabet, getLtrPos(potentialsLtrs[0])-1) // Increments total length by the length of the new ascii letter
				if args[0] == "standard" {
					totalLength++
				}
				c = totalLength - 1
				potentialsLtrs = []rune{}
				ltrCol = 0
				continue
			}
			if len(potentialsLtrs) < 1 {
				// Oh shit, here we go again
				log.Fatal("No letter found after '" + result + "'...")
			}
			ltrCol++
		}
		var candidats = getIntersec(shortLetters, potentialsLtrs)
		if len(candidats) == 1 {
			result += string(candidats[0]) // Adds last letter if not already done
		}
	}
	if doDebug {
		fmt.Print("\n") // Prints only if debug mode is enabled
	}
	fmt.Println(result)
}

func debug(a ...interface{}) {
	if doDebug {
		fmt.Println(a)
	}
}

func getIntersec(a []rune, b []rune) []rune { // Return intersection(s) between 2 rune slices
	var intersec []rune
	for _, r1 := range a {
		for _, r2 := range b {
			if r1 == r2 {
				intersec = append(intersec, r1)
			}
		}
	}
	return intersec
}

func countLines(r io.Reader) (int, error) {
	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}
	for {
		c, err := r.Read(buf)
		count += bytes.Count(buf[:c], lineSep) // Counts occurences of carriage return in the file's buffer

		switch {
		case err == io.EOF:
			return count, nil

		case err != nil:
			return count, err
		}
	}
}

func getLtrPos(ltr rune) int { // Returns corresponding line number
	if ltr > 31 && ltr < 127 {
		return (int(ltr-32) * 9) + 2
	}
	return 0
}

func getMaxLength(s *string, l int) int {
	var maxLength int
	for i := l; i < l+8; i++ {
		r := strings.NewReader(*s) // Create new reader from given string
		line := readLine(r, i)
		if line == "LINE NOT FOUND" {
			return -1
		}
		if len(line) > maxLength {
			maxLength = len(line)
		}
	}
	return maxLength
}

func getPossibilities(ltrCol int, col [8]rune, ltrs *[]rune) {
	if len(*ltrs) == 0 { // Checks every letter in the alphabet
		for l := 1; l < alphaLength-8; l += 9 {
			alphaCol, err := readCol(alphabet, ltrCol, l)
			if err != nil {
				log.Fatal(err)
			}
			if reflect.DeepEqual(alphaCol, col) { // Checks if these cols are trully equals
				*ltrs = append(*ltrs, getRuneAt(l))
			}
		}
	} else {
		for i, r := range *ltrs { // Checks only the given letters
			if i > len(*ltrs)-1 {
				break
			}
			runePos := (int(r-32) * 9) + 1
			alphaCol, err := readCol(alphabet, ltrCol, runePos)
			if err != nil {
				log.Fatal(err)
			}
			if !reflect.DeepEqual(alphaCol, col) {
				*ltrs = removeFromSlice(*ltrs, i) // Removes it from the list if it does not match
			}
		}
	}
}

func getRuneAt(l int) rune {
	runeNbr := (l / 9) + 32 // Gets the rune from the line num
	if runeNbr < 32 || runeNbr > 126 {
		log.Fatal("Invalid rune")
	}
	return rune(runeNbr)
}

func getShortLetters(width int) []rune {
	var ltrs []rune
	for char := ' '; char <= '~'; char++ { // Checks every letter in the 'alphabet'
		if getMaxLength(&alphabet, getLtrPos(char)) <= width {
			ltrs = append(ltrs, char) // Adds it to the list if the width is 3 or less
		}
	}
	return ltrs
}

func isColEmpty(col [8]rune) bool { // Checks if the given column is trully empty
	for _, r := range col {
		if r != 32 {
			return false
		}
	}
	return true
}

func isSpaceFrom(content string, c int, p int) bool {
	for ci := c; ci < c+5; ci++ {
		col, err := readCol(content, ci, p) // Reads a 5 columns sequence
		if err != nil {
			log.Fatal(err)
		}
		if !isColEmpty(col) {
			return false
		}
	}
	return true
}

func readCol(from string, colNum, startline int) ([8]rune, error) {
	var col [8]rune
	for l := startline; l < startline+8; l++ {
		line := readLine(strings.NewReader(from), l)
		if line == "LINE NOT FOUND" {
			if l < startline+7 {
				line = ""
			} else {
				return [8]rune{}, errors.New("line " + strconv.Itoa(l) + " not found")
			}
		}
		if len(line) <= colNum {
			col[l-startline] = ' ' // Adds extra space if the line is not long enough
		} else {
			col[l-startline] = rune(line[colNum])
		}
	}
	return col, nil
}

func readLine(r *strings.Reader, lineNum int) string {
	sc := bufio.NewScanner(r)
	var lastLine int
	for sc.Scan() { // Looks for the good line
		if lastLine == lineNum {
			return sc.Text()
		}
		lastLine++
	}
	return "LINE NOT FOUND"
}

func removeFromSlice(s []rune, i int) []rune {
	if len(s) == 0 {
		return s
	}
	s[len(s)-1], s[i] = s[i], s[len(s)-1] // Moves wanted element to the end
	return s[:len(s)-1]                   // Selects the whole line except the last element
}
