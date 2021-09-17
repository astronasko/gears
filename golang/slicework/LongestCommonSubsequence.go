import "reflect"

// Returns the longest common subsequence of two slices. Based on:
// https://en.wikipedia.org/wiki/Longest_common_subsequence_problem
// Initially written for strings, not tested with slice of integers.
func LongestCommonSubsequence(s1, s2 []int) []int {
	// Corner cases
	if reflect.DeepEqual(s1, s2) {
		return s1
	}

	if len(s1) == 0 || len(s2) == 0 {
		return []int{}
	}
	
	// First Wikipedia property of LCS
	if reflect.DeepEqual(s1[len(s1)-1], s2[len(s2)-1]) {
		return LCS(s1[:len(s1)-1], append(s2[:len(s2)-1], s1[len(s1)-1]...))
	} else {
		// Second Wikipedia property of LCS  
		lcs1 := LCS(s1[:len(s1)-1], s2)
		lcs2 := LCS(s1, s2[:len(s2)-1])
		
		if len(lcs1) > len(lcs2) {
			return lcs1
		} else {
			return lcs2
		}
	}
  }