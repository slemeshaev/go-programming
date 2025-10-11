// Exercise 1.4. Modify dup2 to print the names of all files in which each duplicated line occurs.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	filesWithDups := make(map[string]bool)
	files := os.Args[1:]

	if len(files) == 0 {
		hasDup, err := hasDuplicates(os.Stdin)
		if err != nil {
			fmt.Fprintf(os.Stderr, "reading error stdin:%v\n", err)
			return
		}

		if hasDup {
			fmt.Println("stdin")
		}
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)

			hasDup, err := hasDuplicates(f)
			if err != nil {
				fmt.Fprintf(os.Stderr, "ex1_4: %v\n", err)
				continue
			}

			if hasDup {
				filesWithDups[arg] = true
			}

			f.Close()
		}
	}

	for filename := range filesWithDups {
		fmt.Println(filename)
	}
}

func hasDuplicates(f *os.File) (bool, error) {
	counts := make(map[string]int)
	input := bufio.NewScanner(f)

	for input.Scan() {
		line := input.Text()
		counts[line]++
		if counts[line] > 1 {
			return true, nil
		}
	}

	if err := input.Err(); err != nil {
		return false, err
	}

	return false, nil
}
