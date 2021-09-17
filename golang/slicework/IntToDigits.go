// Decomposes an integer to a slice of digits, ignoring sign.
func IntToDigits(n int) []int {
	var digits []int
	for n > 0 {
		digits = append(digits, n%10)
		n /= 10
	}

	return Revert(digits)
}