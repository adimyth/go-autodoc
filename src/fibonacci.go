package src

/*
Get nth element in the Fibonacci sequence

Parameters
----------
n : int

	The position for which we want the element in the fibonacci sequence

Returns
-------
int

	The nth element in the sequence

Examples
--------
These are written in doctest format, and should illustrate how to
use the function.

```go
n := 10
fmt.Println(FibonacciRecursion(10))
```
*/
func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}
