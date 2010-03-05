package main

import (
	"fmt"
	"./password_creator"
	"flag"
)

var (
	passwordLength = flag.Int("length", 8, "Length of the generated Password")

	//Define all available flags, which are used to specific, which characters to use to generate the password
	lowerCase         = flag.Bool("lower", false, "Should LowerCase characters be included?")
	upperCase         = flag.Bool("upper", false, "Should UpperCase characters be included?")
	numbers           = flag.Bool("numbers", false, "Should the Numbers be included?")
	specialCharacters = flag.Bool("special", false, "Should special characters be included?")
)


func main() {
	flag.Parse()

	password_creator.CreateCharacterArray(*lowerCase, *upperCase, *numbers, *specialCharacters)

	password, error := password_creator.GeneratePassword(*passwordLength)

	if error == nil {
		//Everything went well
		fmt.Println("Your password: ", password)
	} else {
		fmt.Println("Error: ", error.String())
	}
}
