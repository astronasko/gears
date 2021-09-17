// Computes the next lexicographical permutation of a slice of integers. If
// there is no higher lexicographical permutation, returns nil. Based on:
// https://www.nayuki.io/page/next-lexicographical-permutation-algorithm
func NextLexPermutation(a []int) ([]int, nil) {
	// Get length of slice
	n := len(a)

	// Determine length of longest non-increasing suffix
	suf_len := 1

	for i := n-1; i > 0; i-- {
		if a[i-1] < a[i] {
			break
		}
		suf_len += 1
	}

	// Get pivot index
	pivot := n-1-suf_len

	// If pivot is not within bounds, no greater perm!
	if pivot < 0 {
		return nil
	}

	// Find the smallest element rightwards from pivot, yet bigger than pivot!
	successor := pivot+1

	for i := pivot+2; i < n; i++ {
		if (a[pivot] < a[i]) && (a[i] <= a[successor]) {
		successor = i
		}
	}

	// Swap pivot and successor
	a[pivot], a[successor] = a[successor], a[pivot]

	// Reverse slice rightwards from pivot
	b := Revert(a[pivot+1:])
	a = append(a[:pivot+1], b...)

	return a
}