package main

import (
	"fmt"
	"log"
	"path/filepath"
)

func main() {
	files, err := filepath.Glob("*.go")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(files)
}
