/*Copyright (C) 2010 Andreas Sinz

This file is part of GoPasswordCreator.

GoPasswordCreator is free software; you can redistribute it and/or modify it under the terms of the GNU General Public License
as published by the Free Software Foundation; only version 2 of the License.
This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY;
without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.
You should have received a copy of the GNU General Public License along with this program; if not, see <http://www.gnu.org/licenses/>.
*/


package main

import (
	"fmt"
	"flag"
	"os"
	"strings"
)

var (
	passwordLength = flag.Int("length", 8, "Length of the generated Password")

	// Variables that define what characters to use in the password
	lowerCase bool
	upperCase bool
	numerals bool
	specialCharacters bool
	usersCharacters string

	//The user can determine how many passwords will be created
	passwordCount = flag.Int("count", 1, "Determine how many passwords will be created")

	//The user can define a file where the passwords will be written into.
	//If file is omitted, then it will print the passwords on Stdout.
	file = flag.String("file", "", "The File where the passwords should be written into")
)


func usage() {
	command := os.Args[0]
  fmt.Fprintf(os.Stderr,
		`Usage: %s [all] [lower] [upper] [numbers] [special] [own=CHARACTERS]
%s requires at least one of the following commands:
  all: Use lower/upper-case letters, numbers, special characters, and user defined characters to generate the password
  lower: Use lower-case letters
  upper: Use upper-case letters
  numbers: Use digits
  special: Use special characters
  own: Characters defined by the user which will be also be used to generate the password
Options:
`,
		command, command)
  flag.PrintDefaults()
}


func main() {
	flag.Usage = usage
	flag.Parse()

	for _, arg := range flag.Args() {

		// Separate the subcommand from the value
		parsed := strings.SplitN(arg, "=", 2)

		switch parsed[0] {
		case "all":
			lowerCase = true
			upperCase = true
			numerals = true
			specialCharacters = true
		case "lower":
			lowerCase = true
		case "upper":
			upperCase = true
		case "numbers":
			numerals = true
		case "special":
			specialCharacters = true
		case "own":
			if len(parsed) == 2 {
				usersCharacters = parsed[1]
			} else {
				printError(fmt.Errorf("'own' requires a '=' to specify characters"))
			}
		default:
			printError(fmt.Errorf("Invalid argument: %s", parsed[0]))
		}
	}

	var output *os.File
	var fileErr error

	if *file != "" {
		if output, fileErr = os.Create(*file); fileErr != nil {
			printError(fileErr)
			output = os.Stdout
		}
	} else {
		output = os.Stdout
	}

	creator, err := NewCreator(output, lowerCase, upperCase, numerals, specialCharacters, usersCharacters)
	defer output.Close()

	if err != nil {
		printError(err)
	} else {
		writeErr := creator.WritePasswords(*passwordLength, *passwordCount)

		if writeErr != nil  {
			printError(writeErr)
		}
	}
}

func printError(err error) {
	fmt.Println("Error: " + err.Error())
}
