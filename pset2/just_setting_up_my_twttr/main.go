package main

import (
	"fmt"
	"helpers"
//	"strings"
)

func main(){
	user_input := helpers.Input("Input: ")

	target_letters := []string{"a","e","i","o","u","A","E","I","O","U"}
	var final_string string
	count := 0
	for _,char := range user_input{
		if contains(target_letters, string(char)){
			count ++
			continue
		}
		final_string = fmt.Sprintf("%v%v",final_string,string(char))
		count ++
	}
	fmt.Println("Output:",final_string)
}

func contains(list []string, element string) bool {
	for _,e := range list{
		if e == element{
			return true
		}
	}
	return false
}


