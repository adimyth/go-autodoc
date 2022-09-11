package src

import (
	"fmt"
	"math"
)

/*
Calculate square root of a number

Given a float, calculate it's square root. Return error if a negative number is passed.

Parameters
----------
n : float64

	The number to get the square root of.

Returns
-------
float64

	The square root of n.

Raises
------
error

	Error in case a negative number is passed

Examples
--------
These are written in doctest format, and should illustrate how to
use the function.

```go
n := 4.0
res, err := squareRoot(n)

	if err != nil {
		fmt.Print(err)
	}

fmt.Printf("square root of %f: %f", n, res)
```
*/
func SquareRoot(n float64) (float64, error) {
	if n < 0 {
		return n, fmt.Errorf("cannot compute square root of negative numberz")
	}
	return math.Sqrt(n), nil
}
