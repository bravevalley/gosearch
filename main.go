package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var (
	fileMatched []string
)

// Todo
// Complete after learning Go Os and Fs methods.

func fileSearch(root, token string)  {
	cwd, err := os.ReadDir(root)
	if err != nil {
		fmt.Errorf("Unable to open directory - %v", root)
	}

	for _, files := range cwd {
		if strings.Contains(files.Name(), token) {
			fileMatched = append(fileMatched, filepath.Join(root, files.Name()))
		}

		if files.IsDir() {
			fileSearch(files.Name(), token)
		}
	}
}

func main() {
	fileSearch("/", "Goboids")

	for _, v := range fileMatched {
		fmt.Println(v)
	}

}
