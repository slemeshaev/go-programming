// Exercise 1.4. Modify dup2 to print the names of all files in which each duplicated line occurs.
package main

import (
	"bufio"
	"fmt"
	"os"
)

type FileData struct {
	counts map[string]int
	files  map[string][]string
}

func NewFileData() *FileData {
	return &FileData{
		counts: make(map[string]int),
		files:  make(map[string][]string),
	}
}

func (data *FileData) findDuplicates(f *os.File, filename string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		data.counts[line]++

		exists := false
		for _, existingFile := range data.files[line] {
			if existingFile == filename {
				exists = true
				break
			}
		}

		if !exists {
			data.files[line] = append(data.files[line], filename)
		}
	}
}

func main() {
	fileData := NewFileData()
	files := os.Args[1:]

	if len(files) == 0 {
		fileData.findDuplicates(os.Stdin, "stdin")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "ex1_4: %v\n", err)
				continue
			}
			fileData.findDuplicates(f, arg)
			f.Close()
		}
	}

	for line, n := range fileData.counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%v\n", n, line, fileData.files[line])
		}
	}
}
