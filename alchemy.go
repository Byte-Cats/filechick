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

func GenerateLotteryNumbers(num int) []int {
	lotteryNumbers := make([]int, 0, num)
	for len(lotteryNumbers) < num {
		n := rand.Intn(num) + 1
		if !contains(lotteryNumbers, n) {
			lotteryNumbers = append(lotteryNumbers, n)
		}
	}
	return lotteryNumbers
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
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

