// Computes the GCD of two numbers using Euclid's algorithm.
func GreatestCommonDivisor(a, b int) int {
	for b != 0 {
		temp := b
		b = a % b
		a = temp
	}

	return a
}