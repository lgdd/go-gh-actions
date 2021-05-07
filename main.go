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
	if len(args) == 1 {
		cmd := args[0]
		if cmd == "version" {
			fmt.Printf("Version: %s\n", version)
			fmt.Printf("Commit: %s\n", commit)
			fmt.Printf("Date: %s\n", date)
		} else if cmd == "update" {
			doSelfUpdate()
		} else if cmd == "check" {
			latest, found, err := selfupdate.DetectLatest("lgdd/go-gh-actions")
			if err != nil {
				log.Println("Error occurred while detecting version:", err)
				os.Exit(1)
			}
			if found {
				fmt.Printf("Latest version found: \t%v\n", latest.Version)
				fmt.Printf("Current version: \t%v\n", version)
				fmt.Printf("Latest:\n%v\n", latest)
			} else {
				fmt.Println("Latest version not found")
			}
		} else {
			fmt.Println("unknown command")
			os.Exit(1)
		}
	} else if len(args) == 0 {
		fmt.Println("try one of the commands:")
		fmt.Println("- version")
		fmt.Println("- update")
	} else {
		if args[0] == "echo" {
			fmt.Printf("%v", args[1:])
		} else {
			fmt.Println("unknown command")
			os.Exit(1)
		}
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
		log.Printf("Current binary (%v) is the latest version (%v)\n", version, latest.Version)
	} else {
		log.Println("Successfully updated to version", latest.Version)
		log.Println("Release note:\n", latest.ReleaseNotes)
	}
}
