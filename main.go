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
	// If no argument is given print an error and return
	if len(args) != 2 {
		fmt.Println("Please provide a file path")
		return
	}

	ConvertFileToBase64(args[1])
}

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
	var fileSize int64 = fileInfo.Size()
	bytes := make([]byte, fileSize)

	// Read file content to bytes
	_, err = file.Read(bytes)
	if err != nil {
		log.Fatal(err)
		return
	}

	// Encode to base64
	encoded := base64.StdEncoding.EncodeToString(bytes)

	// Create new file
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
