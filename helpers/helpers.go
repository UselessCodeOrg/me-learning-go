package helpers

import (
	"os"
	"fmt"
	"bufio"
	"strings"
)


func Input(text string) string{
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(text)

	user_input,_ := reader.ReadString('\n')

	return strings.TrimSpace(user_input)
}
