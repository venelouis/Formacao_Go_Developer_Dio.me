// https://pkg.go.dev/os#example-OpenFile

package main

import (
	"log" // Package log implements a simple logging package
	"os"  // Package os provides a portable way of using operating system dependent functionality.
)

func main() {
	f, err := os.OpenFile("notes.txt", os.O_RDWR|os.O_CREATE, 0755) // OpenFile opens the named file with specified flag
	if err != nil {                                                 // If there is an error, print it and exit
		log.Fatal(err) // log.Fatal is equivalent to Print() followed by a call to os.Exit(1).
	}
	if err := f.Close(); err != nil { // Close closes the File, rendering it unusable for I/O.
		log.Fatal(err)
	}
}
