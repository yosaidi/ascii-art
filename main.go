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
	input := os.Args[2]
	input = strings.ReplaceAll(input, `\n`, "\n")

	runes := []rune(input)
	isValid := ascii.AreStringValid(runes)
	if !isValid {
		fmt.Println("invalid character in your input  ")
		return
	}

	filename:=""
	font:="standard"
	
		if strings.HasPrefix(os.Args[1], "--output=") &&
			strings.HasSuffix(os.Args[1], ".txt") {
				filename = os.Args[1][9:]
		
		if os.Args[3] == "standard"||  os.Args[3]=="thinkertoy" || os.Args[3] == "shadow" {
			font = os.Args[3]
		}

	}

	

	ascii.Print(&input, filename,font)

}
