// Exercise 4.9: Write a program wordfreq to report the frequency of each word
// in an input text file. Call input.Split(bufio.ScanWords) before the first call
// to Scan to break the input into words instead of lines.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: wordfreq <filename>\n")
		os.Exit(1)
	}

	filename := os.Args[1]

	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "wordfreq: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	freq := make(map[string]int)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		word := scanner.Text()
		word = strings.ToLower(word)
		word = strings.Trim(word, ".,!?;:()[]{}'\"`")

		if word != "" {
			freq[word]++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "wordfreq: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Word frequencies (case-insensitive):")
	fmt.Println("-----------------------------------")

	words := make([]string, 0, len(freq))
	for word := range freq {
		words = append(words, word)
	}

	// sorting words
	for i := 0; i < len(words)-1; i++ {
		for j := i + 1; j < len(words); j++ {
			if words[i] > words[j] {
				words[i], words[j] = words[j], words[i]
			}
		}
	}

	// printing
	for _, word := range words {
		fmt.Printf("%-20s %d\n", word, freq[word])
	}

	fmt.Println("-----------------------------------")
	fmt.Printf("Total unique words: %d\n", len(freq))

	total := 0
	for _, count := range freq {
		total += count
	}
	fmt.Printf("Total words: %d\n", total)
}
