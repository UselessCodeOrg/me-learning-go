package main

import (
	"os"
	"fmt"
	"bufio"
	"strconv"
	"strings"
)

func input(text string) string{

	fmt.Print(text)

	reader := bufio.NewReader(os.Stdin)
	user_input,_ := reader.ReadString('\n')

	return strings.TrimSpace(user_input)
}

func dollar_to_float(text string) float64{

	i,_ := strconv.Atoi(text)
	f := float64(i)
	return f
}

func percent_to_float(text string) float64{

	i,_ := strconv.Atoi(text)
	f := float64(i)
	result := f/100
	return result
}

func main(){
	dollars := dollar_to_float(input("Cost of meal: "))
	percent := percent_to_float(input("Tip percent: "))
	tip := dollars*percent
	fmt.Println("Tip to pay: ",tip)
}
