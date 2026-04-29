// Exercise 4.11: Build a tool that lets users create, read, update, and delete GitHub issues from the command line,
// invoking their preferred text editor when substantial text input is required.

package main

import (
	"fmt"
	"log"
	"os"
)

var usage = "%s Usage:\n\tsearch QUERY\nOr:\n\t[read|edit|close|open] OWNER REPO ISSUE_NUMBER\n"

func usageDie() {
	fmt.Fprintf(os.Stderr, usage, os.Args[0])
	os.Exit(1)
}

func main() {
	if len(os.Args) < 2 {
		usageDie()
	}

	cmd := os.Args[1]
	args := os.Args[2:]

	if cmd == "search" {
		if len(args) < 1 {
			usageDie()
		}
		search(args)
		os.Exit(0)
	}

	if len(args) != 3 {
		usageDie()
	}

	owner, repo, number := args[0], args[1], args[2]
	switch cmd {
	case "read":
		readIssue(owner, repo, number)
	case "edit":
		editIssue(owner, repo, number)
	case "close":
		closeIssue(owner, repo, number)
	case "open":
		openIssue(owner, repo, number)
	}
}

func search(query []string) {
	result, err := SearchIssues(query)
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range result.Items {
		format := "#%-6d\t%20.20s\t%.100s\n"
		fmt.Printf(format, item.Number, item.User.Login, item.Title)
	}
}

func readIssue(owner, repo, number string) {
	// Implementation
}

func editIssue(owner, repo, number string) {
	// Implementation
}

func openIssue(owner, repo, number string) {
	// Implementation
}

func closeIssue(owner, repo, number string) {
	// Implementation
}
