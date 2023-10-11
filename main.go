package main

import (
	"crypto/sha512"
	"fmt"
	"io"
	"os"

	"github.com/barasher/go-exiftool"
)

func main() {
	et, err := exiftool.NewExiftool()
	if err != nil {
		fmt.Printf("Error when intializing: %v\n", err)
		return
	}
	defer et.Close()

	fileInfos := et.ExtractMetadata("test.jpeg")

	for _, fileInfo := range fileInfos {
		if fileInfo.Err != nil {
			fmt.Printf("Error concerning %v: %v\n", fileInfo.File, fileInfo.Err)
			continue
		}

		for k, v := range fileInfo.Fields {
			fmt.Printf("[%v] %v\n", k, v)
		}
	}

	sha512, err := sha512sum("test.jpeg")
	if err != nil {
		fmt.Printf("Error when calculating sha512: %v\n", err)
		return
	}
	fmt.Println("[SHA512] ", sha512)
}

func sha512sum(filePath string) (string, error) {
	fileHandle, err := os.Open(filePath)
	if err != nil {
		return "", err
	}

	defer fileHandle.Close()

	sha512 := sha512.New()
	if _, err := io.Copy(sha512, fileHandle); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", sha512.Sum(nil)), nil
}
