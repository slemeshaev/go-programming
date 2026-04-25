// Exercise 4.10: Modify issues to report the results in age categories, say less than a month old,
// less than a year old, and more than a year old.

package main

import (
	"chapter4/cgithub/cgithub"
	"fmt"
	"log"
	"os"
	"time"
)

const (
	lessThanMonth = iota
	lessThanYear
	moreThanYear
)

var categories = []string{
	"Less than a month old",
	"Less than a year old",
	"More than a year old",
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: issues <query>\n")
		os.Exit(1)
	}

	result, err := cgithub.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	issues := make([][]*cgithub.Issue, 3)
	now := time.Now()

	for _, item := range result.Items {
		age := now.Sub(item.CreatedAt)

		switch {
		case age < 30*24*time.Hour:
			issues[lessThanMonth] = append(issues[lessThanMonth], item)
		case age < 365*24*time.Hour:
			issues[lessThanYear] = append(issues[lessThanYear], item)
		default:
			issues[moreThanYear] = append(issues[moreThanYear], item)
		}
	}

	for i, category := range categories {
		if len(issues[i]) == 0 {
			fmt.Printf("\n=== %s ===\nNo issues found.\n", category)
			continue
		}

		fmt.Printf("\n=== %s ===\n", category)
		for _, issue := range issues[i] {
			age := time.Since(issue.CreatedAt)
			fmt.Printf("#%-5d %10.0f days ago %9s %s\n", issue.Number, age.Hours()/24, issue.User.Login, issue.Title)
		}
	}
}
