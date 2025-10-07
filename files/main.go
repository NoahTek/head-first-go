package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func scanDirectory(path string) error {
	fmt.Println("│\n📁", path)
	files, err := os.ReadDir(path)
	if err != nil {
		return err
	}
	for _, file := range files {
		filePath := filepath.Join(path, file.Name())
		if file.IsDir() {
			err := scanDirectory(filePath)
			if err != nil {
				return err
			}
		} else {
			fmt.Println("└──🗎", filePath)
		}
	}
	return nil
}

func main() {
	err := scanDirectory("getfloat")
	if err != nil {
		log.Fatal(err)
	}
}
