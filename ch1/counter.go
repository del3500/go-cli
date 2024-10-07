package main

import (
	"bufio"
	"io"
	"log"
)

// A scanner is a convenient way of reading data delimited
// by spaces or new lines.

func count(r io.Reader, countWords, countBytes bool) int {
	scanner := bufio.NewScanner(r)
	if countBytes && countWords {
		log.Fatal("cannot use both -w (words) and -b (bytes) at the same time")
	}

	if countBytes {
		buf := make([]byte, 4096)
		totalBytes := 0
		for {
			n, err := r.Read(buf)
			if n > 0 {
				totalBytes += n
			}
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal(err)
			}
		}
		return totalBytes
	}

	switch {
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
