package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func reportPanic() {
	p := recover()
	if p == nil {
		return
	}
	err, ok := p.(error)
	if ok {
		fmt.Println(err)
	} else {
		panic(p) // for unanticipated panics or if type assertions fails.
	}
}

func scanDirectory(path string) {
	fmt.Println("│\n📁", path)
	files, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		filePath := filepath.Join(path, file.Name())
		if file.IsDir() {
			scanDirectory(filePath)
		} else {
			fmt.Println("└──🗎", filePath)
		}
	}
}

func main() {
	defer reportPanic()
	// simulating a different panic scenario where type assertion fails.
	// panic("some other issue")
	scanDirectory("getfloat")
}
