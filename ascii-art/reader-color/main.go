package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"runtime"
	"strconv"
	"strings"
)

var glines [8]string
var alphabet string
var color string
var background string
var index int
var custom bool
var liste_maxLenth []int
var balise bool
var liste_balise []int
var interval int
var begin int
var head int
var tail int
var isInterval bool
var effect string

func main() {
	args := os.Args[1:]
	if len(args) < 1 || len(args[0]) < 1 { // Check if there is only one argument
		return
	}
	text := args[0]

	SetStringFlag(&color, "color", "white", &args)
	SetStringFlag(&background, "background", "", &args)
	setFlag()

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
		if carRtrns[i] { // If there is a cariage return:
			printColorLines()    // Displays the characters that have already been added to the lines
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

	printColorLines()
}

func getLtrPos(ltr rune) int { // Returns corresponding line number
	if ltr > 31 && ltr < 127 {
		return (int(ltr-32) * 9) + 2
	}
	return 0
}

func printColorLines() {

	if balise {
		setInterval(head, tail)
	}

	for j, line := range glines { // Prints each line
		for i, c := range line { // Prints each char of the current line
			if color == "rainbow" {
				Rainbow(j, c)
			} else if background == "cocorico" {
				cocorico(i, c)
			} else {
				if isInterval {
					if i <= begin || i > interval {
						fmt.Print(string(c))
					} else {
						fmt.Print("\u001b["+effect+"38;2;"+color+background+"m", string(c), "\u001b[0m")
					}
				} else {
					fmt.Print("\u001b["+effect+"38;2;"+color+background+"m", string(c), "\u001b[0m")
				}
			}
		}
		fmt.Print("\n")

	}
}

func cocorico(i int, c rune) {
	if i <= (len(glines[0]) / 3) {

		fmt.Print("\u001b["+effect+"38;2;"+color+";48;2;0;128;255"+"m", string(c), "\u001b[0m")

	} else if i <= (len(glines[0])/3)*2 {
		fmt.Print("\u001b["+effect+"38;2;"+color+";48;2;255;255;255"+"m", string(c), "\u001b[0m")

	} else {

		fmt.Print("\u001b["+effect+"38;2;"+color+";48;2;255;00;00"+"m", string(c), "\u001b[0m")
	}
}

func setInterval(a, b int) {

	isInterval = true
	begin = 0
	interval = 0

	for i := 0; i < a; i++ {
		begin += liste_maxLenth[i] + 1

	}
	begin--
	interval += begin

	for i := a; i < b; i++ {
		interval += liste_maxLenth[i] + 1
	}

}

func setFlag() {

	color = strings.ToLower(color)
	background = strings.ToLower(background)

	if color[len(color)-1] == '-' {
		color = color[:len(color)-1]
	}

	if strings.Contains(color, "-bold") {
		effect += "1;"
		index := strings.Index(color, "-bold")
		color = color[:index] + color[index+len("-bold"):]
	}

	if strings.Contains(color, "-italic") {
		effect += "3;"
		index := strings.Index(color, "-italic")
		color = color[:index] + color[index+len("-italic"):]
	}

	if strings.Contains(color, "-underline") {
		effect += "4;"
		index := strings.Index(color, "-underline")
		color = color[:index] + color[index+len("-underline"):]
	}

	if strings.Contains(color, "-flash") {
		effect += "5;"
		index := strings.Index(color, "-flash")

		color = color[:index] + color[index+len("-flash"):]
	}
	if strings.Contains(color, "[") {
		balise = true

		index := strings.Index(color, "[")
		index2 := strings.Index(color, ",")
		index3 := strings.Index(color, "]")

		head, _ = strconv.Atoi(string(color[index+1 : index2]))
		tail, _ = strconv.Atoi(string(color[index2+1 : index3]))

		color = color[:index]
	}

	if strings.Contains(color, "-") {
		color = color[:strings.Index(color, "-")]
	}

	if len(color) < 1 || color == "white" {
		color = "255;255;255"
	} else if color[0] == '#' && color != "rainbow" {
		custom = true
		color = GetHexaColor(color[1:3]) + ";" + GetHexaColor(color[3:5]) + ";" + GetHexaColor(color[5:7])
	} else if color != "White" && !custom && color != "rainbow" {
		switch color {
		case "black":
			color = "00;00;00"
		case "red":
			color = "255;00;00"
		case "green":
			color = "77;255;00"
		case "yellow":
			color = "255;255;00"
		case "brown":
			color = "130;90;30"
		case "blue":
			color = "0;128;255"
		case "grey":
			color = "150;150;150"
		case "orange":
			color = "255;182;00"
		case "purple":
			color = "150;00;255"
		case "pink":
			color = "255;00;255"
		default:
			color = "255;255;255"
		}
	} else if color != "rainbow" {
		color = "255;255;255"
	}

	if len(background) > 0 {
		if background != "" && background[0] != '#' && background != "cocorico" {
			switch background {
			case "black":
				background = ";48;2;00;00;00"
			case "red":
				background = ";48;2;255;00;00"
			case "green":
				background = ";48;2;0;114;41"
			case "yellow":
				background = ";48;2;255;255;00"
			case "brown":
				background = ";48;2;130;90;30"
			case "blue":
				background = ";48;2;0;128;255"
			case "grey":
				background = ";48;2;150;150;150"
			case "orange":
				background = ";48;2;255;182;00"
			case "purple":
				background = ";48;2;150;00;255"
			case "pink":
				background = ";48;2;255;00;255"
			default:
				background = ";48;2;255;255;255"
			}
		} else if background[0] == '#' && background != "cocorico" {
			background = ";48;2;" + GetHexaColor(background[1:3]) + ";" + GetHexaColor(background[3:5]) + ";" + GetHexaColor(background[5:7])
		}
		fmt.Println(background)
		if background[0] != ';' && background != "cocorico" {
			background = "0"
		}
	}
}

func GetHexaColor(hexa string) string {

	var liste []int

	for _, valstr := range hexa {
		for i, valbase := range "0123456789ABCDEF" {
			if valstr == valbase {
				liste = append(liste, i)
			}
		}
	}

	number := liste[0]
	for _, val := range liste[1:] {
		number = number*16 + val
	}
	return strconv.Itoa(number)

}

func Rainbow(i int, c rune) {

	listecolor := [6]string{"255;00;00", "255;182;00", "255;255;00", "10;90;229", "150;00;255", "255;00;255"}

	color := i % 6

	fmt.Print("\u001b["+effect+"38;2;"+listecolor[color]+background+"m", string(c), "\u001b[0m")
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
	liste_maxLenth = append(liste_maxLenth, maxLength)
}

func SetStringFlag(variable *string, name, defaultValue string, args *[]string) {
	if len(*args) > 1 {
		for ai, arg := range *args {
			if strings.HasPrefix(arg, "--"+name+"=") { // Checks if the current arg begins with the given param name
				sep := -1
				for i, c := range arg {
					if c == '=' {
						sep = i // Gets the position of the start of the value
						break
					}
				}
				if sep < 0 {
					break
				}
				value := arg[sep+1:] // Gets the value out of the argument
				*variable = value    // Sets given variable value to the argument value
				argsCopy := *args
				before := argsCopy[:ai]          //
				after := argsCopy[ai+1:]         // Remove argument from args list
				*args = append(before, after...) //
				return
			}
		}
	}
	*variable = defaultValue // Sets given variable value to default value
}
