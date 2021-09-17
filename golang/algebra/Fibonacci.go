// Computes n-th Fibonacci member iteratively.
// Time O(n), Memory O(1)
func Fibonacci(n int) int {
	a, b := 0, 1

	for i := 0; i < n ; i++ {
		a, b = b, a+b
	}

	return a
}

// Computes n-th Fibonacci number recursively.
// Time O(φ^n), Memory O(n)
// http://courses.csail.mit.edu/6.01/spring07/lectures/lecture4.pdf
func FibonacciRecursive(n int) int {
	if n < 2 {
		return n
	}

	return Fibonacci(n-1) + Fibonacci(n-2)
}