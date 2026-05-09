package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

func input(text string) string {
	fmt.Print(text)
	reader := bufio.NewReader(os.Stdin)
	user_input,_ := reader.ReadString('\n')
	return user_input
}

func convert(text string) string {
	nstring_1 := strings.Replace(text,":)","🙂",-1)
	nstring_2 := strings.Replace(nstring_1,":(","🙁",-1)
	return nstring_2
}

func main(){
	user_input := input("Input: ")
	fmt.Println("Converted string: ",convert(user_input))
}
