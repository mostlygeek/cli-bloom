package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/willf/bloom"
)

func main() {
	flag.Parse()
	m := flag.Arg(0)
	if m == "" {
		fmt.Println("Usage:", os.Args[0], "<max size>")
		os.Exit(1)
	}

	mN, err := strconv.ParseUint(m, 10, 64)
	if err != nil {
		fmt.Println("Error:", err.Error())
	}

	_ = mN

	bFilter := bloom.New(uint(mN)*10, 5)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		b := scanner.Bytes()
		if bFilter.TestAndAdd(b) {
			fmt.Printf("Dup: %s\n", string(b))
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
