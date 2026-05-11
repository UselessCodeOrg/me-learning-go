package main

import (
	"fmt"
	"helpers"
	"strings"
	"strconv"
)

func main(){
	user_input := helpers.Input("What time is it? ")
	time := convert(user_input)
	if time > 7 && time < 8{
		fmt.Println("breakfast time")
	}
        if time > 12 && time < 13{
                fmt.Println("lunch time")
        }
        if time > 18 && time < 19{
                fmt.Println("dinner time")
        }

}

func convert(text string) float64{
	s_text := strings.Split(text,":")

	hrs_sti,_ := strconv.Atoi(s_text[0])
	mins_sti,_ := strconv.Atoi(s_text[1])

	hrs_f := float64(hrs_sti)
	mins_f := float64(mins_sti)

	decimal := mins_f/60
	final := hrs_f+decimal

	return final
}
