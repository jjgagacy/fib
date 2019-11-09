package fib_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/jjgagacy/fib"
)

func TestFib(t *testing.T) {
	for i := 0; i < 9; i++ {
		fmt.Print(strconv.Itoa(fib.FibonacciLoop(i)) + " ")
	}
	fmt.Println("")
	for i := 0; i < 9; i++ {
		fmt.Print(strconv.Itoa(fib.FibonacciRecursion(i)) + " ")
	}
	fmt.Println()

	fmt.Println(fib.Fib(5))
}
