package main

import (
	"strings"
)

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
				*variable = value // Sets given variable value to the argument value
				argsCopy := *args
				before := argsCopy[:ai]  //
				after := argsCopy[ai+1:]     // Remove argument from args list
				*args = append(before, after...) //
				return
			}
		}
	}
	*variable = defaultValue // Sets given variable value to default value
}
