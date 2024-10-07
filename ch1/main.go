package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var words bool
	var bytes bool
	flag.BoolVar(&bytes, "b", false, "count of bytes")
	flag.BoolVar(&words, "w", false, "count of words")
	fmt.Println(count(os.Stdin, words, bytes))
}
