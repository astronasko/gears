// Count number of "ones" in the binary representation of a positive integer.
func CountBits(input int) int {
	ones := 0
	// Exploit that in binary arithmetic, division by two moves
	// the entire principal part rightwards (to the decimal pt)
	// So check how many times a number is odd (i.e. ends with 1)
	// to count number of bits, and then move rightwards
	for input != 0 {
		if input % 2 == 1 {
		ones +=1
		}
		input /= 2
	}
	
	return ones
  }