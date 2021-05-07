package main

import (
	"fmt"
	"log"
	"os"

	"github.com/blang/semver"
	"github.com/rhysd/go-github-selfupdate/selfupdate"
)

var (
	version string
	commit  string
	date    string
)

func main() {
	args := os.Args[1:]
	fmt.Printf("%v\n", args)
	if len(args) == 1 {
		cmd := args[0]
		if cmd == "version" {
			fmt.Printf("Version: %s\n", version)
			fmt.Printf("Commit: %s\n", commit)
			fmt.Printf("Date: %s\n", date)
		} else if cmd == "update" {
			doSelfUpdate()
		} else {
			fmt.Println("unknown command")
			os.Exit(1)
		}
	} else if len(args) == 0 {
		fmt.Println("try one of the commands:")
		fmt.Println("- version")
		fmt.Println("- update")
	} else {
		fmt.Println("unknown command")
		os.Exit(1)
	}
}

func doSelfUpdate() {
	v := semver.MustParse(version)
	latest, err := selfupdate.UpdateSelf(v, "lgdd/go-gh-actions")
	if err != nil {
		log.Println("Binary update failed:", err)
		os.Exit(1)
	}
	if latest.Version.Equals(v) {
		log.Println("Current binary is the latest version", version)
	} else {
		log.Println("Successfully updated to version", latest.Version)
		log.Println("Release note:\n", latest.ReleaseNotes)
	}
}
