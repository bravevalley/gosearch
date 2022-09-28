package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// Todo
// Complete after learning Go Os and Fs methods.

var (
	fileMatches []string
)

func searchFile(searchDir, token string) {
	fmt.Println("Started!!!")
	cwd := os.DirFS(searchDir)
	fmt.Print(cwd.Open("var"))
	listOfFiles, _ := fs.ReadDir(cwd, "/")

	for _, files := range listOfFiles {
		if strings.Contains(files.Name(), token) {
			fileMatches = append(fileMatches, filepath.Join(searchDir, files.Name()))
		}

		if files.IsDir() {
			searchFile(files.Name(), token)
		}
	}
}

func main() {
	searchFile("/", "go")
	for _, v := range fileMatches {
		fmt.Printf("Found : %s", v)
	}

}
