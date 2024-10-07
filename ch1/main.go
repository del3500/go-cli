package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	var bytes bool
	var words bool
	flag.BoolVar(&words, "w", false, "count lines")
	flag.BoolVar(&bytes, "b", false, "count bytes")
	flag.Parse()
	fmt.Println(count(os.Stdin, words, bytes))
}

// A scanner is a convenient way of reading data delimited
// by spaces or new lines.

func count(r io.Reader, countWords, countBytes bool) int {
	scanner := bufio.NewScanner(r)
	if countBytes && countWords {
		log.Fatal("cannot use both -w (words) and -b (bytes) at the same time")
	}
	switch {
	case countBytes:
		scanner.Split(bufio.ScanBytes)
	case countWords:
		scanner.Split(bufio.ScanWords)
	default:
		scanner.Split(bufio.ScanLines)
	}
	wc := 0
	for scanner.Scan() {
		wc++
	}
	return wc
}
