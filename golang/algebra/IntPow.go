// Raise a integer to a non-negative integer power. Replaces the need of using
// math.Pow for elementary exponentiation.
func IntPow(a,b int) int {
	for b > 1 {
		a *= a
		b -= 1
	}
	return a
}