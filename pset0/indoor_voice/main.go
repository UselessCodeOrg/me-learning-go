package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Input something: ")

	userInput, _ := reader.ReadString('\n')

	lower := strings.ToLower(strings.TrimSpace(userInput))

	fmt.Println(lower)
}
