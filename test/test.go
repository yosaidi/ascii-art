package main

import (
	"fmt"
	"os"
	"strings"
)


func main(){
input := os.Args[1]
result := strings.Split(input,"\n")
for i := 0; i < len(result); i++ {
	fmt.Println(result[i])
}
}