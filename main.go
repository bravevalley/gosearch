package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

var (
	fileMatched []string
	wg          = sync.WaitGroup{}
	mu          = sync.Mutex{}
)

// File search function that takes the search token and working dir as an arg
func fileSearch(root, token string) {

	// Reduce the wait group counter by one
	defer wg.Done()

	// ioutils is deprecated, so this is how we read Dir
	// The Working Dir string is convert to a FS and queried
	cwd, err := os.ReadDir(root)

	// UX just to let the user know whats going on
	fmt.Printf("Reading from %v\n", root)

	// Catches error when unable to open a Dir
	if err != nil {
		fmt.Printf("Unable to open directory - %v\n", root)
	}

	// This loops over the files found in the wordking Dir
	for _, files := range cwd {

		// Test each files if the name of the file contains the search token
		if strings.Contains(files.Name(), token) {

			// Locked the memory the thread is writing to, to prevent race
			//  conditions
			mu.Lock()

			// Updates the slice declared with the files whose name contains
			//  the search token
			fileMatched = append(fileMatched, filepath.Join(root, files.Name()))

			// Unlocks the memory so other threads can secure the lock and
			//  write to it
			mu.Unlock()
		}

		// Test if the file is a Dir so it can cd into the Dir and continue its
		//  search
		if files.IsDir() {

			// Because we are calling a recursive function, the defered wg.Done
			//  () is called thereby reducing the waitgroup counter and
			//    bringing the code to an end. So we need to increase the
			//     waitgroup counter at every instance so the main wg.done is
			//      not uttered
			wg.Add(1)

			// Create another goroutine for the recursion
			go fileSearch(filepath.Join(root, files.Name()), token)
		}
	}
}

func main() {

	// Seed the time the program start
	start := time.Now()

	// Increment the waitgroup before creating a goroutine
	wg.Add(1)

	// Spin up the GoRoutine
	go fileSearch("/home/dassyareg/Downloads", ".pem")

	// Suspend execution till the WaitGroup counter equals zero
	wg.Wait()

	// Print the a format text of the files matched
	for _, v := range fileMatched {
		fmt.Printf("Matched : %5v\n", v)
	}

	// Print time elapsed
	fmt.Println(time.Since(start).Seconds())
}
