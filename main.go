package main

import (
	"context"
	"fmt"

	"github.com/barasher/go-exiftool"
	"github.com/luizhreis/file-metadata/filehandler"
	"github.com/luizhreis/file-metadata/hasher"
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

	fileHandlerService := filehandler.New()

	fileHandler, err := fileHandlerService.OpenFile(context.TODO(), filePath)

	if err != nil {
		panic(err)
	}

	hasherService := hasher.New()

	hashes, err := hasherService.GetHashes(context.TODO(), fileHandler)
	if err != nil {
		panic(err)
	}

	fmt.Println(hashes)

	fileHandler.Close()
}
