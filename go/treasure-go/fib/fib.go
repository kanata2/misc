package main

import (
	"flag"
	"fmt"
	"os"
)

var n int

func main() {
	flag.IntVar(&n, "n", 0, "argument of fib function")
	flag.Parse()

	os.Exit(run())
}

func run() int {
	if n < 0 {
		fmt.Fprintf(os.Stderr, "n is not negative number: %d\n", n)
		return 1
	}
	fmt.Printf("fib(%d) = %d\n", n, fib(n))
	return 0
}

func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
