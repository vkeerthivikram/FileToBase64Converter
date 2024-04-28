package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Please provide a file path")
		return
	}

	ConvertFileToBase64(args[1])
}

// ConvertFileToBase64 converts a file to its Base64 representation and saves it as a new file.
// The function takes a file path as input and reads the file content into a byte array.
// It then encodes the byte array to Base64 format and creates a new file with the encoded data.
// The new file is saved in the same directory as the original file, with a name starting with "base64_",
// followed by the original file name and a timestamp in Unix format.
// The function returns an error if any of the file operations fail.
func ConvertFileToBase64(fp string) {
	file, err := os.Open(fp)

	if err != nil {
		log.Fatal(err)
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
		log.Fatal(err)
		return
	}
	encoded := base64.StdEncoding.EncodeToString(bytes)
	timestamp := time.Now().Unix()
	baseFileName := filepath.Base(fp)
	newFileName := fmt.Sprintf("base64_%s_%d.txt", baseFileName, timestamp)
	newFilePath := filepath.Join(filepath.Dir(fp), newFileName)
	err = os.WriteFile(newFilePath, []byte(encoded), 0644)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("File successfully encrypted to base64 and saved as %s\n", newFileName)
}
