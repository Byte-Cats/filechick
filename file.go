package filechick

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"

	"github.com/spf13/afero"
	"github.com/spf13/viper"
)

var Afero = afero.Afero{Fs: afero.NewOsFs()}

// CreateEmptyFile creates an empty file.
// Returns the file pointer and an error if any.
func CreateEmptyFile(fileName string) (*os.File, error) {
	file, err := os.Create(fileName)
	if err != nil {
		return nil, err
	}
	return file, nil
}

// DeleteFile deletes the specified file.
// Returns an error if the operation fails.
func DeleteFile(file string) error {
	err := os.RemoveAll(file)
	if err != nil {
		return fmt.Errorf("error deleting file: %w", err)
	}
	return nil
}

// SaveHtml saves HTML content from a URL to a specified file.
// Returns an error if the operation fails.
func SaveHtml(url string, fileName string) error {
	// Validate URL
	if _, err := regexp.Compile(url); err != nil { // Changed from http.ParseRequestURI to regexp.Compile
		return fmt.Errorf("invalid URL: %w", err)
	}

	// Create the file
	file := CreateFile(fileName) // Changed to only assign the file
	if file == nil {
		return fmt.Errorf("error creating file") // Updated error handling
	}
	defer file.Close()

	// Get HTML content
	res, reqErr := CustomRequest(url)
	if reqErr != nil {
		return fmt.Errorf("error getting HTML from URL: %w", reqErr)
	}

	// Write to file
	if _, fileErr := file.WriteString(res); fileErr != nil {
		return fmt.Errorf("error writing to file: %w", fileErr)
	}

	return nil
}

// CreateFile creates a file.
// Returns the file pointer and an error if any.
func CreateFile(fileName string) *os.File {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating file", err)
		return nil
	}
	return file
}

// LoadHtml loads HTML content from a file to a string.
// Returns the HTML content as a string and an error if any.
func LoadHtml(file string) (string, error) {
	// Open the file
	f, err := os.Open(file)
	if err != nil {
		return "", fmt.Errorf("error opening file: %w", err)
	}
	defer f.Close()

	// Read the file
	bytes, err2 := io.ReadAll(f)
	if err2 != nil {
		return "", fmt.Errorf("error reading file: %w", err2)
	}

	// Turn result into a string
	return string(bytes), nil
}

// NewDir creates a new directory.
// Returns an error if the operation fails.
func NewDir(dir string) error {
	// If directory doesn't exist, create it
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.Mkdir(dir, 0755)
		if err != nil {
			return fmt.Errorf("error creating directory: %w", err)
		}
	} else if err != nil {
		return fmt.Errorf("error creating directory: %w", err)
	}
	return nil
}

// ExitIfExists exits the program if the file already exists.
func ExitIfExists(dir string) {
	if _, err := os.Stat(dir); !os.IsNotExist(err) {
		fmt.Println("Filechick detected that you already have this file downloaded. Exiting...")
		os.Exit(0)
	}
}

// SaveImage saves an image from a URL to a file.
// Returns an error if the operation fails.
func SaveImage(url string, filename string) error {
	// Get the image
	results, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error getting image: %w", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Error closing body:", err)
		}
	}(results.Body)

	// Create an empty file
	emptyFile, err := CreateEmptyFile(filename)
	if err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}
	defer emptyFile.Close()

	// Copy the image to the file
	_, copyErr := io.Copy(emptyFile, results.Body)
	if copyErr != nil {
		return fmt.Errorf("error copying file: %w", copyErr)
	}

	return nil
}

// TitleToDirName converts a title to a directory name.
// Returns the directory name as a string.
func TitleToDirName(title string) string {
	reg, _ := regexp.Compile("[^a-zA-Z\\d]+")
	return reg.ReplaceAllString(title, "")
}

// RemoveIfExists removes a file if it exists.
// Returns an error if the operation fails.
func RemoveIfExists(path string) error {
	exists, err := Afero.Exists(path)

	if err != nil {
		return fmt.Errorf("error checking file existence: %w", err)
	}

	if exists {
		err = Afero.Remove(path)
		if err != nil {
			return fmt.Errorf("error removing file: %w", err)
		}
	}
	return nil
}

// GetFileNames returns a slice of strings containing all the file names in a directory.
// Returns an error if the operation fails.
func GetFileNames(dir string) ([]string, error) {
	files, err := Afero.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("error reading directory: %w", err)
	}
	var fileNames []string
	for _, file := range files {
		fileNames = append(fileNames, file.Name())
	}
	return fileNames, nil
}

// FileOrDirExists checks if a file or directory exists.
// Returns true if the file or directory exists, false otherwise.
func FileOrDirExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// CopyFile copies a file from the source to the destination.
// Returns an error if the operation fails.
func CopyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("error opening source file: %w", err)
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("error creating destination file: %w", err)
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return fmt.Errorf("error copying file: %w", err)
	}

	return out.Close()
}

// ReadFileLineByLine reads a file line by line.
// Returns a slice of strings containing the file lines and an error if any.
func ReadFileLineByLine(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error scanning file: %w", err)
	}
	return lines, nil
}

// VippyEnv gets an environment variable value using viper.
// Returns the value as a string.
func VippyEnv(key string) string {
	// Use viper to get the value from the environment variable
	vippy := viper.New()
	vippy.SetConfigName(".env")
	vippy.SetConfigType("env")
	vippy.AddConfigPath(".")
	vippy.AllowEmptyEnv(false)
	vippy.AutomaticEnv()
	err := vippy.ReadInConfig()
	if err != nil {
		log.Fatal("Error reading env file: ", err)
	}
	return vippy.GetString(key)
}
