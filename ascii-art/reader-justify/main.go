package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

var glines [8]string
var alphabet string

func main() {
	var err error
	var output []byte

	args := os.Args[1:]
	if len(args) < 1 || len(args) > 4 || len(args[0]) < 1 { // Check if there is only one argument
		return
	}
	text := args[0]

	width, _ := getTermSize() // Gets the width and the heigth of the terminal
	// fmt.Printf("%d rows by %d cols\n", heigth, width)

	var align string
	SetStringFlag(&align, "align", "left", &args) // Gets argument 'align' if present; if not set its value to defaultValue
	// fmt.Println("Alignment:", align)

	if len(args) == 1 {
		args = append(args, "standard")
	}

	_, currentFilePath, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("Failed to find current dir")
	}
	currentDir := path.Dir(currentFilePath) // Get current dir path

	if len(args) < 2 {
		args = append(args, "shadow")
	}
	if args[1] == "shadow" {
		output, err = ioutil.ReadFile(path.Join(currentDir, "..", "fonts", "shadow.txt")) // Read shadow.txt file
	} else if args[1] == "thinkertoy" {
		output, err = ioutil.ReadFile(path.Join(currentDir, "..", "fonts", "thinkertoy.txt")) // Read thinkertoy.txt file
	} else {
		output, err = ioutil.ReadFile(path.Join(currentDir, "..", "fonts", "standard.txt")) // Read standard.txt file
	}

	if err != nil {
		panic(err)
	}
	alphabet = string(output) // Stores the string conversion of the content in `alphabet`

	for i, ltr := range text { // Iterates through given text
		suffixLine(ltr)      // Adds current letter line by line to lines variables (stored in `glines`)
		if i < len(text)-1 { // Adds little separation between letters
			for l := range glines {
				glines[l] += " "
			}
		}
	}

	var padding int
	if align == "center" {
		padding = (width - len(glines[0])) / 2 // Sets the left padding to half the width minus the len of the final ascii text
	} else if align == "right" {
		padding = width - len(glines[0]) // Sets the left padding to the width minus the len of the final ascii text
	} else if align == "justify" {
		if len(args) > 1 {
			if args[1] == "shadow" { // Adapts the width to the font
				width--
			} else if args[1] == "thinkertoy" {
				width -= 2
			}
		}
		remaining := width - len(glines[0])
		reg := regexp.MustCompile(" ")
		spaceNbr := len(reg.FindAllStringIndex(text, -1)) // Counts the number of spaces
		if spaceNbr > 0 {
			spaceWidth := remaining / spaceNbr
			// fmt.Println("SpaceNbr:", spaceNbr, "| SpaceWidth:", spaceWidth)
			spacePos := getSpacesPos() // Gets the positions of the spaces in the final ascii text
			var newSpace string
			for i := 0; i < spaceWidth; i++ { // Creates the new width of a space
				newSpace += " "
			}
			// fmt.Println("SpacePos:", spacePos)
			for _, sp := range spacePos {
				insertSpaces(newSpace, sp) // Inserts new spaces at every position referred in `spacePos`
			}
		}
	}
	// fmt.Println(width, padding, len(glines[0]))
	for _, line := range glines { // Prints each line one after the other
		for i := 0; i < padding; i++ {
			line = " " + line
		}
		for _, c := range line { // Prints each char of the current line
			fmt.Print(string(c))
		}
		fmt.Print("\n")
	}
}

func insertSpaces(space string, pos int) {
	for i := range glines {
		glines[i] = glines[i][:pos] + space + glines[i][pos:] // Inserts new space at the specified position
	}
}

func getLtrPos(ltr rune) int { // Returns corresponding line number
	if ltr > 31 && ltr < 127 {
		return (int(ltr-32) * 9) + 2
	}
	return 0
}

func getSpacesPos() []int {
	var total []int
	var followV, followH int
	for i := range glines[0] { // Checks if every line at this index is really a space
		for j := range glines {
			char := glines[j][i]
			if char != ' ' {
				followV = 0
				followH = 0
				break
			}
			followV++
		}
		if followV > 7 {
			followH++
		}
		if followH > 5 {
			total = append([]int{i}, total...) // Adds the position to the final list if there is a space of at least 4 wide
			followH = 0
		}
	}
	return total
}

func getTermSize() (int, int) {
	var width, heigth int
	var errW, errH error
	goos := runtime.GOOS
	if goos == "linux" || goos == "darwin" { // UNIX systems
		cmd := exec.Command("stty", "size") // Executes the 'stty' commands, which returns informations about the current console
		cmd.Stdin = os.Stdin
		out, err := cmd.Output()
		if err != nil {
			log.Fatal(err)
		}
		var space int
		sOut := string(out)
		for i, c := range sOut { // Gets the position of the space
			if c == ' ' {
				space = i
				break
			}
		}
		width, errW = strconv.Atoi(sOut[:space]) // Outputs the width from the command's result
		if errW != nil {
			log.Fatal(errW)
		}
		heigth, errH = strconv.Atoi(sOut[space+1 : len(sOut)-1]) // Outputs the height from the command's result
		if errH != nil {
			log.Fatal(errH)
		}
	} else if goos == "windows" { // Windows systems
		cmd := exec.Command("mode", "con") // Executes the 'mode' commands, which returns informations about the current console
		cmd.Stdin = os.Stdin
		out, err := cmd.Output()
		if err != nil {
			log.Fatal(err)
		}
		rows := strings.Split(string(out), "\n") // Parses command output
		row := rows[4][:len(rows[4])-1]
		lastSpace := strings.LastIndex(row, " ")
		cols, err := strconv.Atoi(row[lastSpace+1:]) // Turns cols stats to a number
		if err != nil {
			log.Fatal(err)
		}
		return cols, 0
	} else {
		fmt.Println("Sorry, but your computer does not have a classic operating system")
		os.Exit(1)
	}
	return heigth, width
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
