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
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	passwordLength = 8

	// Variables that define what characters to use in the password
	lowerCase         bool
	upperCase         bool
	numerals          bool
	specialCharacters bool
	usersCharacters   string

	// Make sure there will always be at least one password
	passwordCount = 1

	file string
)

func usage() {
	command := os.Args[0]
	fmt.Fprintf(os.Stderr,
		`Usage: %s [all] [alphanum] [lower] [upper] [numbers] [special] [own=CHARACTERS]
%s requires at least one of the following subcommands to specify what characters
may be used in the password:
  all: Equivalent to 'alphanum special'
  alphanum: Equivalent to 'lower upper numbers'
  lower: Use lower-case letters
  upper: Use upper-case letters
  numbers: Use digits
  special: Use special characters
  own: Specifies a custom set of characters to use
'all', 'alphanum', 'lower', 'upper', 'numbers', and 'special' may be followed by
'=f' to nullify that character set.
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

		// Group arguments by the data type of their value
		switch parsed[0] {
		case "own":
			// Need a string value
			if len(parsed) == 2 {
				usersCharacters = parsed[1]
			} else {
				printError(fmt.Errorf("'own' requires a '=' to specify characters"))
			}
		case "-count":
			// Need a int value
			if len(parsed) == 2 {
				// Fetch the int value from parsed[1]
				var err error
				passwordCount, err = strconv.Atoi(parsed[1])
				if err != nil {
					printError(err)
				}
			} else {
				printError(fmt.Errorf("'-count' requires a '=' to specify number of passwords"))
			}
		case "-length":
			// Need a int value
			if len(parsed) == 2 {
				// Fetch the in value from parsed[1]
				var err error
				passwordLength, err = strconv.Atoi(parsed[1])
				if err != nil {
					printError(err)
				}
			} else {
				printError(fmt.Errorf("'-length' requires a '=' to specify the length passwords should be"))
			}
		case "-file":
			// Need a string value
			if len(parsed) == 2 {
				file = parsed[1]
			} else {
				printError(fmt.Errorf("'-file' requires a '=' to specify the file name to write passwords to"))
			}
		default:
			// All other arguments take boolean values
			on := true
			if len(parsed) == 2 {
				var err error
				on, err = strconv.ParseBool(parsed[1])
				if err != nil {
					printError(err)
				}
			}
			switch parsed[0] {
			case "all":
				lowerCase = on
				upperCase = on
				numerals = on
				specialCharacters = on
			case "alphanum":
				lowerCase = on
				upperCase = on
				numerals = on
			case "lower":
				lowerCase = on
			case "upper":
				upperCase = on
			case "numbers":
				numerals = on
			case "special":
				specialCharacters = on
			default:
				printError(fmt.Errorf("Invalid argument: %s", parsed[0]))
			}
		}

	}

	var output *os.File
	var fileErr error

	if file != "" {
		if output, fileErr = os.Create(file); fileErr != nil {
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
		writeErr := creator.WritePasswords(passwordLength, passwordCount)

		if writeErr != nil {
			printError(writeErr)
		}
	}
}

func printError(err error) {
	fmt.Fprintln(os.Stderr, "Error: "+err.Error())
}
