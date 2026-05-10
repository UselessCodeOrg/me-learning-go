package main


import (
	"os"
	"fmt"
	"bufio"
	"strings"
)

func input(text string) string{
	fmt.Print(text)
	reader := bufio.NewReader(os.Stdin)
	user_input,_ := reader.ReadString('\n')
	return strings.TrimSpace(user_input)
}

func main(){
	user_input := input("Greeting: ")
	l_user_input := strings.ToLower(user_input)
	if strings.HasPrefix(l_user_input,"h"){
		if strings.HasPrefix(l_user_input,"hello"){
			fmt.Println("$0")
		} else {
			fmt.Println("$20")
		}
	} else {
		fmt.Println("$100")
	}
}
