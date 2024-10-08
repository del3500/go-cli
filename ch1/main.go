package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	var words bool
	var bytes bool
	flag.BoolVar(&bytes, "b", false, "count of bytes")
	flag.BoolVar(&words, "w", false, "count of words")
	flag.Parse()
	t := strings.NewReader("hello")
	fmt.Println(count(t, words, bytes))
}
