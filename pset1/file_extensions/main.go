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
	data := map[string]string{
		"gif":"image",
		"jpg":"image",
		"jpeg":"image",
		"png":"image",
		"pdf":"application",
		"txt":"text",
		"zip":"application",
	}
	user_input := input("File extension: ")
	s_user_input := strings.Split(user_input,".")
	if len(s_user_input) == 1{
		return
	}

	ext := s_user_input[1]
	_,ok := data[ext]
	if !ok {
		fmt.Println("application/octet-stream")
		return
	}

	fmt.Printf("%s/%v\n",data[ext],ext)
}
