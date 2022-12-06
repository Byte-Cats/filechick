package filechick

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// GetInput Function to get user input from the command line
func GetInput() string {
	var input string
	in := bufio.NewReader(os.Stdin)

	input, err := in.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return ""
	}
	input = strings.Replace(input, "\n", "", -1)

	return input
}

