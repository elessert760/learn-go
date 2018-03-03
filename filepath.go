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
    fmt.Println(len(files)) // contains a list of all files in the current directory
}
