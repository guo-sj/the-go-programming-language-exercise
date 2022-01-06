
// Exercise 1.3: Experiment to measure the difference in running time between our potentially
// inefficient versions and the one that uses strings.Join. (Section 1.6 illustrates part of 
// the time package, and Section 11.4 shows how to write benchmark tests for systematic per-
// formance evaluation.)

// Conclusion:
// inefficient version: 23.789µs
// strings.Join       :  4.106µs

package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	var s, sep string

	// inefficient version
	start := time.Now()
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	t := time.Now()
	fmt.Println("inefficient version: ", t.Sub(start))

	// strings.Join version
	start = time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	t = time.Now()
	fmt.Println("strings.Join version: ", t.Sub(start))
}
