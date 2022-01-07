
// Exercise 1.8: Modify 'fetch' to add the prefix 'http://' to each argument URL if it is missing.
// You might want to use 'strings.HasPrefix'

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	const prefix = "http://"
	for _, url := range os.Args[1:] {
		if strings.HasPrefix(url, prefix) == false {
			url = prefix + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stdout, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stdout, "fetch: reading %s: %v\n", b, err)
			os.Exit(1)
		}
		fmt.Printf("%s\n", b)
	}
}
