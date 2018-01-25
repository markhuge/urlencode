package main

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
)

func main() {
	bytes, err := ioutil.ReadAll(os.Stdin)
	// Quick n dirty
	if err != nil {
		panic(err)
	}
	// Get rid of that %0A at the end
	bytes = bytes[:len(bytes)-1]

	out := url.PathEscape(string(bytes))

	fmt.Println(out)
}
