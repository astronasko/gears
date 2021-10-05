// Decomposes an integer into a map of [primes]powers. For example,
// PrimeFactorisation(4030) == map[2:1 5:1 13:1 31:1], as 4030 == 2*5*13*31.
func PrimeFactorisation(n int) fact_map map[int]int {
	// Start with dividing by 2
	for n%2 == 0 {
		fact_map[2] += 1
		n /= 2
	}
	// Now use higher integers with step of 2
	for i := 3; i <= n; i += 2 {
		for n%i == 0 {
			fact_map[i] += 1
			n /= i
		}
	}

	return
}
// Can be further optimised, but I haven't had timeouts with it for now.