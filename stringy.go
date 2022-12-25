package filechick

import (
	"strconv"
	"strings"
	"unicode"
	"time"
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
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// RemoveWords removes words from a string
func RemoveWords(s string, words []string) string {
	for _, word := range words {
		s = strings.Replace(s, word, "", -1)
	}
	return s
}

// PrintCharMessage prints a given string character by character with a given delay time between each character.
// This is useful for displaying a message in an animated way. 
// message: The string message that needs to be printed.
// delay: The delay time between each character in milliseconds. 
func PrintCharMessage(message string, delay time.Duration) {
    for _, character := range message {
        fmt.Printf("\r%s", string(character))
        time.Sleep(delay)
    }
    fmt.Println()
}

// RemoveWord removes a word from a string
func RemoveWord(s string, word string) string {
	return strings.Replace(s, word, "", -1)
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

// RemoveWhitespaceAndPunctuation removes whitespace and punctuation from a string
func RemoveWhitespaceAndPunctuation(s string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) || unicode.IsPunct(r) {
			return -1
		}
		return r
	}, s)
}

// RemoveWhitespaceAndPunctuationAndNumbers removes whitespace, punctuation, and numbers from a string
func RemoveWhitespaceAndPunctuationAndNumbers(s string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) || unicode.IsPunct(r) || unicode.IsNumber(r) {
			return -1
		}
		return r
	}, s)
}

// RemoveNumbers removes numbers from a string
func RemoveNumbers(s string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsNumber(r) {
			return -1
		}
		return r
	}, s)
}

// IsNumber checks if a string is a number
func IsNumber(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

// IsInt checks if a string is an integer
func IsInt(s string) bool {
	_, err := strconv.ParseInt(s, 10, 64)
	return err == nil
}

// StringToFloat converts a string to a float
func StringToFloat(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}

// FloatToString converts a float to a string
func FloatToString(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

// StringToBool converts a string to a bool
func StringToBool(s string) (bool, error) {
	return strconv.ParseBool(s)
}

// BoolToString converts a bool to a string
func BoolToString(b bool) string {
	return strconv.FormatBool(b)
}

// StringToRune converts a string to a rune
func StringToRune(s string) rune {
	return []rune(s)[0]
}

// RuneToString converts a rune to a string
func RuneToString(r rune) string {
	return string(r)
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
