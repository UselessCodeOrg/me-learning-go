package main

import (
	"fmt"
	"helpers"
	"strings"
	"strconv"
)

func main(){

	user_input := helpers.Input("Expression: ")

	s := strings.Split(user_input," ")
	x,y,z := s[0],s[1],s[2]

	x_i,_ := strconv.Atoi(x)
	z_i,_ := strconv.Atoi(z)

	if y == "+" {
		fmt.Printf("%.1f\n",float64(x_i+z_i))
	}
        if y == "-" {
                fmt.Printf("%.1f\n",float64(x_i-z_i))
        }
        if y == "*" {
                fmt.Printf("%.1f\n",float64(x_i*z_i))
        }
        if y == "/" {
                fmt.Printf("%.1f\n",float64(x_i/z_i))
        }

}
