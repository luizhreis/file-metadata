package main

import (
	"crypto/sha256"
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

	filePath := "test.jpeg"

	fileInfos := et.ExtractMetadata(filePath)

	for _, fileInfo := range fileInfos {
		if fileInfo.Err != nil {
			fmt.Printf("Error concerning %v: %v\n", fileInfo.File, fileInfo.Err)
			continue
		}

		for k, v := range fileInfo.Fields {
			fmt.Printf("[%v] %v\n", k, v)
		}
	}

	sha512, err := sha512sum(filePath)
	if err != nil {
		fmt.Printf("Error when calculating sha512: %v\n", err)
		return
	}
	fmt.Println("[SHA512] ", sha512)

	sha256, err := sha256sum(filePath)
	if err != nil {
		fmt.Printf("Error when calculating sha256: %v\n", err)
		return
	}
	fmt.Println("[SHA256] ", sha256)
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

func sha256sum(filePath string) (string, error) {
	fileHandle, err := os.Open(filePath)
	if err != nil {
		return "", err
	}

	defer fileHandle.Close()

	sha256 := sha256.New()
	if _, err := io.Copy(sha256, fileHandle); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", sha256.Sum(nil)), nil
}
