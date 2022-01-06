// Exercise 1.2: Modify the echo program to print the index and value of each of its arguments,
// one per line.

package main

import (
	"fmt"
	"os"
)

func main() {
	var sep string

	sep = ": "
	for i := 0; i < len(os.Args); i++ {
		fmt.Println(i, sep, os.Args[i])
	}
}
