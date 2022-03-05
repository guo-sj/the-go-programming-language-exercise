// Exercise 4.3: Rewrite reverse to use an array pointer instead of a slice.

package main

import "fmt"

const max = 10

func main() {
	s := &[max]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	reverse(s)
	fmt.Println(*s)
}

func reverse(s *[max]int) {
	for i, j := 0, len(*s)-1; i < j; i, j = i+1, j-1 {
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
	}
}
