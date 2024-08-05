package filechick

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"unicode"
)

// StringToInt converts a string to an int
func StringToInt(str string) (int, error) {
	return strconv.Atoi(str)
}

// IntToString converts an int to a string
func IntToString(i int) string {
	return strconv.Itoa(i)
}

// ReverseString reverses a string
func ReverseString(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// RemoveWords removes words from a string
func RemoveWords(s string, words []string) string {
	for _, word := range words {
		s = strings.ReplaceAll(s, word, "")
	}
	return s
}

// PrintCharMessage prints a given string character by character with a given delay time between each character.
func PrintCharMessage(message string, delay time.Duration) {
	for _, character := range message {
		fmt.Printf("\r%s", string(character))
		time.Sleep(delay)
	}
	fmt.Println()
}

// RemoveWhitespace removes whitespace from a string
func RemoveWhitespace(s string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, s)
}

// RemovePunctuation removes punctuation from a string
func RemovePunctuation(s string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsPunct(r) {
			return -1
		}
		return r
	}, s)
}

// ReplaceWhitespaceWithHyphen replaces whitespace with hyphens in a string
func ReplaceWhitespaceWithHyphen(s string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return '-'
		}
		return r
	}, s)
}

// TruncateString truncates a string to a given length
func TruncateString(s string, length int) string {
	if len(s) > length {
		return s[:length]
	}
	return s
}

// BreakUpString breaks up a string into lines of a given length
func BreakUpString(s string, interval int) string {
	words := strings.Fields(s) // Use Fields to split by whitespace
	var newString strings.Builder
	charCount := 0

	for _, word := range words {
		if charCount+len(word) > interval {
			newString.WriteString("\n")
			charCount = 0
		}
		newString.WriteString(word + " ")
		charCount += len(word) + 1
	}

	return newString.String()
}
