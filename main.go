package main

import "fmt"

var (
	version string
	commit  string
	date    string
)

func main() {
	fmt.Printf("Version: %s\n", version)
	fmt.Printf("Commit: %s\n", commit)
	fmt.Printf("Date: %s\n", date)
}
