// Exercise 1.10: Find a web site that produces a large amount of data. Investigate caching by
// running fetchall twice in succession to see whether the reported time changes much. Do
// you get the same content each time? Modify fetchall to prints its output to a file so it can be
// examined.

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string) // create a channel by make utility
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
	}
	const filename = "result"
	f, _ := os.Open(filename)
	for range os.Args[1:] {
		//fmt.Println(<-ch) // receive from channel ch
		fmt.Fprintf(f, "%s\n", <-ch)
	}
	//fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
	fmt.Fprintf(f, "%.2fs elapsed\n", time.Since(start).Seconds())
	f.Close()
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}
