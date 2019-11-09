package fib

// FibonacciLoop returns the nth fibonacci number
//
// Calculating the fibonacci number with slices and loops
func FibonacciLoop(n int) int {
    f := make([]int, n+1, n+2)
    if n < 2 {
        f = f[0:2]
    }
    f[0] = 0
    f[1] = 1
    for i := 2; i <= n; i++ {
        f[i] = f[i-1] + f[i-2]
    }
    return f[n]
}

// FibonacciRecursion returns the nth fibonacci number
//
// Calculating the fibonacci number with recursion
func FibonacciRecursion(n int) int {
    if n <= 1 {
        return n
    }
    return FibonacciRecursion(n - 1) + FibonacciRecursion(n - 2)
}

// Fib calls FibonacciLoop
func Fib(n int) int {
    return FibonacciLoop(n)
}
