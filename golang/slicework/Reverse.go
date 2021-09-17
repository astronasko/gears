// Reverses a slice, i.e. turns a one-dimensional slice around.
func Reverse(a []int) []int {
	n := len(a)

	for i := 0; i < n/2; i++ {
		j := n-i-1
		a[i], a[j] = a[j], a[i]
	}
	
	return a
}