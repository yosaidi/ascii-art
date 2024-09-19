package main

import (
	ascii "ascii/helperfunctions"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
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
	myString := ascii.ReadFile()
	lines := strings.Split(input, "\n")

	for i, line := range lines {
	
		if line == "" {
			if i== len(lines)-1{
				continue
			}
			fmt.Println()
			continue
		}
		newInput := strings.Fields(line)

		for row := 1; row < 9; row++ {
			for _, word := range newInput {
				wordRunes := []rune(word)
				for j := 0; j < len(wordRunes); j++ {
					char := wordRunes[j]
					starts := (int(char) - 32) * 9
					fmt.Print(myString[starts+row])
				}
				fmt.Print(" ")
			}
			fmt.Println()
		}

	}

}
