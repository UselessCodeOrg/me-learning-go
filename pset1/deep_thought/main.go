package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
)

func input(text string) string {

	fmt.Print(text)

	reader := bufio.NewReader(os.Stdin)
	user_input,_ := reader.ReadString('\n')

	return strings.TrimSpace(user_input)
}

func main(){

	user_input := input("What is the answer to the Great Question of Life, The Universe, and Everything? ")
	if user_input == "42" || user_input == "forty-two" || user_input == "forty two" {
		fmt.Println("Yes")
	} else{
		fmt.Println("No")
	}
}
