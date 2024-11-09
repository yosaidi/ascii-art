package ascii

import (
	"fmt"
	"strings"
)

func Print(input *string, font string) {
	str := *input
	lines := strings.Split(str, "\n")
	myString := ReadFile(font)


	for i, line := range lines {
		if line == "" {
			if i == len(lines)-1 {
				continue
			}
			fmt.Println()
			continue
		}
		newInput := SpaceManager(line)

		for row := 1; row < 9; row++ {
			for _, word := range newInput {
				wordRunes := []rune(word)
				for j := 0; j < len(wordRunes); j++ {
					char := wordRunes[j]
					starts := (int(char) - 32) * 9
					fmt.Print(myString[starts+row])
				}
			}
		fmt.Println()
		}
	}

	*input = str
}
