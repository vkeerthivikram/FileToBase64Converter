package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

// main is the entry point of the program. It takes a file path as a command-line argument,
// converts the contents of the file to base64, and writes the result to a new file.
//
// Parameters:
//
//	None
//
// Returns:
//
//	nil - If the operation is successful.
//	error - If an error occurs during file operations.
//
// Example:
//
//	go run main.go input.txt
func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Please provide a file path")
		return
	}

	err := ConvertFileToBase64(args[1])
	if err != nil {
		log.Fatal("error converting file to base64: ", err)
		return
	}
}

// ConvertFileToBase64 reads a file at the specified path, converts its contents to base64,
// and writes the result to a new file. The new file's name is based on the original file's name
// with "_base64_" and a timestamp appended.
//
// Parameters:
//
//	fp (string) - The path to the input file.
//
// Returns:
//
//	nil - If the operation is successful.
//	error - If an error occurs during file operations.
//
// Example:
//
//	ConvertFileToBase64("input.txt")
func ConvertFileToBase64(fp string) error {
	file, err := os.Open(fp)

	if err != nil {
		return err
	}

	defer func(file *os.File) {
		err2 := file.Close()
		if err2 != nil {
			log.Fatal(err2)
		}
	}(file)

	fileInfo, _ := file.Stat()
	var fileSize = fileInfo.Size()
	bytes := make([]byte, fileSize)
	_, err = file.Read(bytes)
	if err != nil {
		return err
	}
	encoded := base64.StdEncoding.EncodeToString(bytes)
	timestamp := time.Now().Unix()
	baseFileName := filepath.Base(fp)
	newFileName := fmt.Sprintf("base64_%s_%d.txt", baseFileName, timestamp)
	newFilePath := filepath.Join(filepath.Dir(fp), newFileName)
	err = os.WriteFile(newFilePath, []byte(encoded), 0644)
	if err != nil {
		return err
	}
	fmt.Printf("File successfully encrypted to base64 and saved as %s\n", newFileName)
	return nil
}
