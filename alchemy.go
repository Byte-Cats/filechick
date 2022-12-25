package filechick

import "github.com/skip2/go-qrcode"

// GenerateQR generates a QR code with the given size, filename and content
func GenerateQR(size int, filename string, content string) {
	// Create a new QR code with the given content
	qr, err := qrcode.New(content, qrcode.Medium)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Write the QR code to the given file
	qr.WriteFile(size, filename)
}

// GenerateLotteryNumbers generates a slice of unique integers from 1 to num.
// num is the number of integers to generate.
// Returns the slice of generated integers.
func GenerateLotteryNumbers(num int) []int {
	// Initialize an empty slice of integers with a capacity of num.
	lotteryNumbers := make([]int, 0, num)

	// Generate num unique integers from 1 to num.
	for len(lotteryNumbers) < num {
		// Generate a random integer from 1 to num.
		n := rand.Intn(num) + 1

		// Check if the generated integer is already in the slice.
		if !contains(lotteryNumbers, n) {
			// If the integer is not in the slice, append it to the slice.
			lotteryNumbers = append(lotteryNumbers, n)
		}
	}

	// Return the slice of generated integers.
	return lotteryNumbers
}

// contains checks if an integer is in a slice of integers.
// s is the slice of integers to check.
// e is the integer to check for.
// Returns true if e is in s, false otherwise.
func contains(s []int, e int) bool {
	// Iterate through the slice of integers.
	for _, a := range s {
		// Check if the current integer is equal to e.
		if a == e {
			// If the integers are equal, return true.
			return true
		}
	}
	// If no integers in the slice are equal to e, return false.
	return false
}

// GeneratePassword generates a random password with the specified length and set of allowed characters.
// length is the length of the generated password.
// chars is a string containing all the characters that are allowed in the generated password.
// Returns the generated password.
func GeneratePassword(length int, chars string) string {
	// Seed the random number generator.
	rand.Seed(time.Now().UnixNano())

	// Initialize a string builder with a capacity of length.
	sb := strings.Builder{}
	sb.Grow(length)

	// Generate length random characters from the chars string.
	for i := 0; i < length; i++ {
		// Generate a random index into the chars string.
		idx := rand.Intn(len(chars))

		// Append the character at the generated index to the string builder.
		sb.WriteByte(chars[idx])
	}

	// Return the resulting string.
	return sb.String()
}

// ConvertPDFToWord converts a PDF file to a Word document.
// inputFilePath is the path of the PDF file.
// outputFilePath is the path of the Word document.
// Returns an error if one occurs during conversion.
func ConvertPDFToWord(inputFilePath string, outputFilePath string) error {
	return pdfcpuAPI.Convert(inputFilePath, outputFilePath)
}

// ConvertWordToPDF converts a Word document to a PDF file.
// inputFilePath is the path of the Word document.
// outputFilePath is the path of the PDF file.
// Returns an error if one occurs during conversion.
func ConvertWordToPDF(inputFilePath string, outputFilePath string) error {
	return pdfcpuAPI.Convert(inputFilePath, outputFilePath)
}

