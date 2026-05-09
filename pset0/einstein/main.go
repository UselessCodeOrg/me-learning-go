package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

func input(text string) string{
	fmt.Print(text)
	reader := bufio.NewReader(os.Stdin)
	user_input,_ := reader.ReadString('\n')
	return strings.TrimSpace(user_input)
}

func main(){

	c := 300000000
	c_s := c*c

	user_input := input("m: ")
	m,_ := strconv.Atoi(user_input)
	fmt.Println("E: ",m*c_s)
}
