
// Exercise 1.4: Modify dup2 to print the names of all files in which each duplicate line occurs.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	stdinCounts := make(map[string]int)
	files := os.Args[1: ]
	if len(files) == 0 {
		countLines(os.Stdin, stdinCounts)
		printLines("Stdin", stdinCounts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			fileCounts := make(map[string]int)    // every time refresh map
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, fileCounts)
			printLines(arg, fileCounts)
			f.Close()
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}

func printLines(arg string, counts map[string]int) {
	i := false
	for line, n := range counts {
		if n > 1 {
			if arg != "Stdin" && i == false {
				fmt.Printf("\n%s:\n", arg)    // print file name
			    i = true;
			}
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
