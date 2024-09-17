package ascii

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFile() []string {
	fileName := "templates/standard.txt"
	input, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer input.Close()
	scanner := bufio.NewScanner(input)
	mySlice := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		mySlice = append(mySlice, line)
	}
	return mySlice
}