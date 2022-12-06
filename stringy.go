package filechick 
import (
	"fmt"
	"strconv"
	"strings"
)

// StringToInt to change string to int
func StringToInt(str string) int {
	var i int
	i, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Error converting string to integer:", err)
		return 0
	}
	return i
}

// FormatString formats the input string to be more readable for a Discord message.
func FormatString(input string) string {
	// Replace newlines with line breaks.
	formatted := strings.ReplaceAll(input, "\n", "\n\n")

	// Add a space after each period.
	formatted = strings.ReplaceAll(formatted, ".", ".\n")

	// Add a space after each comma.
	formatted = strings.ReplaceAll(formatted, ",", ", ")

	return formatted
}
