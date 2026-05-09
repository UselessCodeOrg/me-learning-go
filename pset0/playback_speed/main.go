package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)


func main(){
	//setting up reader
	reader := bufio.NewReader(os.Stdin)

	//taking input and sanitizing
	fmt.Print("Input: ")
	user_input,_ := reader.ReadString('\n')
	trim_uinput := strings.TrimSpace(user_input)

	fmt.Println(strings.Replace(trim_uinput," ","...",-1))
}
