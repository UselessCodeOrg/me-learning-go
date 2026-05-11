package main

import (
	"fmt"
	"helpers"
	"strings"
)

func main(){
	snake_case := ""

	user_input := helpers.Input("camelCase: ")
	lower := strings.ToLower(user_input)
	for i,_ := range user_input{
		if user_input[i] != lower[i]{
			snake_case = fmt.Sprintf("%s_%v",string(snake_case),string(lower[i]))
			continue
		}
		snake_case = fmt.Sprintf("%s%v",string(snake_case),string(user_input[i]))
	}
	fmt.Println("snake_case: ",snake_case)
}
