package filechick

import (
	"fmt"
	"strconv"
	"strings"
)

// StringToInt converts a string to an integer
func StringToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Error converting string to integer:", err)
		return 0
	}
	return i
}

// IntToString converts an integer to a string
func IntToString(i int) string {
	return strconv.Itoa(i)
}

// SplitString splits a string into a slice of strings based on a delimiter
func SplitString(str, delimiter string) []string {
	return strings.Split(str, delimiter)
}

// HasPrefix checks if a string starts with a specific prefix
func HasPrefix(str, prefix string) bool {
	return strings.HasPrefix(str, prefix)
}

// HasSuffix checks if a string ends with a specific suffix
func HasSuffix(str, suffix string) bool {
	return strings.HasSuffix(str, suffix)
}

// Contains checks if a string contains a specific substring
func Contains(str, substr string) bool {
	return strings.Contains(str, substr)
}

// TrimSpace trims leading and trailing whitespace from a string
func TrimSpace(str string) string {
	return strings.TrimSpace(str)
}

// ToLower converts a string to lowercase
func ToLower(str string) string {
	return strings.ToLower(str)
}

// ToUpper converts a string to uppercase
func ToUpper(str string) string {
	return strings.ToUpper(str)
}

// IsNumber checks if a string is a number
func IsNumber(str string) bool {
	_, err := strconv.ParseFloat(str, 64)
	return err == nil
}

// BreakUpString breaks up a string into a slice of strings
func BreakUpString(s string, interval int) string {
	// split the string into words
	words := strings.Split(s, " ")
	// initialize a new string
	newString := ""
	// initialize a counter to track the number of characters in the current line
	charCount := 0

	// loop through the words
	for _, word := range words {
		// if adding the current word to the current line would exceed the interval,
		// add a line break and reset the character count
		if charCount+len(word) > interval {
			newString += "\n"
			charCount = 0
		}

		// add the current word to the new string and increment the character count
		newString += word + " "
		charCount += len(word) + 1
	}

	return newString
}
