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

var alphabet string
var glines [8]string
var out string
var err error

func main() {
	var output []byte
	args := os.Args[1:]

	setStringFlag(&out, "output", "Filename.txt", &args)

	if len(args) > 2 || len(args[0]) < 1 { // Check if there is only one argument
		return
	}

	text := args[0]

	if len(args) == 1 {
		args = append(args, "standard")
	}

	_, currentFilePath, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("Failed to find current dir")
	}
	currentDir := path.Dir(currentFilePath) // Get current dir path
	file := getFile(currentDir)

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
			printLines(file)     // Displays the characters that have already been added to the lines
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
	printLines(file)
	defer file.Close() // Close the file in defer mode
}

func getFile(currentDir string) *os.File {
	var fileExists bool
	_, err := os.Stat(path.Join(currentDir, out)) // Check file

	if err == nil {
		fileExists = true
	}

	for fileExists { // Check if the file exist or not
		fmt.Println("A file with the same name was found, do you want to rename the new file (r), overwrite the old file (o), or cancel ? (c)")
		rep := ""
		_, err = fmt.Scanln(&rep) // Take the first line of the reponse

		if err != nil {
			panic(err)
		}

		if rep == "r" {
			fmt.Println("Enter the new file name")
			_, err = fmt.Scanln(&out) // Take the first line of the input

			if err != nil {
				panic(err)
			}
		} else if rep == "o" {
			_ = os.Remove(path.Join(currentDir, out)) // Remove the older file

		} else {
			fmt.Println("Exiting !")
			os.Exit(0)
		}
		_, err := os.Stat(path.Join(currentDir, out))

		if err == nil {
			fileExists = true
		} else {
			fileExists = false
		}
	}

	emptyFile, err := os.OpenFile(path.Join(currentDir, out), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644) // Create,open the file with write permission only and write the line into the file
	if err != nil {
		log.Fatal(err)
	}
	return emptyFile
}  //Manage file checks

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

func getLtrPos(ltr rune) int { // Returns corresponding line number
	if ltr > 31 && ltr < 127 {
		return (int(ltr-32) * 9) + 2
	}
	return 0
}

func setStringFlag(variable *string, name, defaultValue string, args *[]string) {
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

func printLines(file *os.File) {
	for _, line := range glines {
		_, err = file.WriteString(line + "\n") // Write each line to the file
	}
} //Manage writing