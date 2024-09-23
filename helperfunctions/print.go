package ascii

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Print(input *string, filename,font string) {
	str := *input
	lines := strings.Split(str, "\n")
	myString := ReadFile(font)

	// Create or open the file for writing
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	for i, line := range lines {
		if line == "" {
			if i == len(lines)-1 {
				continue
			}
			writer.WriteString("\n")
			continue
		}
		newInput := SpaceManager(line)

		for row := 1; row < 9; row++ {
			for _, word := range newInput {
				wordRunes := []rune(word)
				for j := 0; j < len(wordRunes); j++ {
					char := wordRunes[j]
					starts := (int(char) - 32) * 9
					writer.WriteString(myString[starts+row])
				}
			}
			writer.WriteString("\n")
		}
	}

	// Flush the buffer to the file
	writer.Flush()
	*input = str
}
