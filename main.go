package main

import (
	ascii "ascii/helperfunctions"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) >4 {
		fmt.Println("invalid args")
		return
	}
	input := os.Args[1]
	input = strings.ReplaceAll(input, `\n`, "\n")

	runes := []rune(input)
	isValid := ascii.AreStringValid(runes)
	if !isValid {
		fmt.Println("invalid character in your input  ")
		return
	}

	font:="standard"
	
	
	ascii.Print(&input,font)

}
