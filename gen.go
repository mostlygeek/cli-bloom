package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	flag.Parse()

	var num, max int64
	var err error

	if flag.NArg() != 2 {
		fmt.Println("Prints random <num> numbers from 0 to <max>")
		fmt.Println("Usage:", os.Args[0], "<num> <max>")
		os.Exit(1)
	} else {
		num, err = strconv.ParseInt(flag.Arg(0), 10, 64)
		if err != nil {
			fmt.Println("Invalid num to generate")
			os.Exit(1)
		}

		max, err = strconv.ParseInt(flag.Arg(1), 10, 64)
		if err != nil {
			fmt.Println("Invalid max to generate")
			os.Exit(1)
		}
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := int64(0); i < num; i++ {
		fmt.Printf("%d\n", r.Int63n(max))
	}

}
