# Golang Fibonacci

[Fibonacci](http://en.wikipedia.org/wiki/Fibonacci_number) is the sequence as follows:

```
1 1 2 3 5 8 13 21 34 55 89 144 233 377 610 …
```

The sequence can start with a zero or a one.

# Usage

fib uses two ways to calculate fibonacci number, `FibonacciLoop` and `FibonacciRecursion`. There is another shortcut method `Fib`, which calls `FibonacciLoop` since `FibonacciRecursion` isn't always efficient.

```go
fmt.Println(fib.Fib(5)) // 5
```
