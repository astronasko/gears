// Remove given argument from given slice.
func ArgDel(slice []int, arg int) []int {
    slice[arg] = slice[len(slice)-1]
    return s[:len(slice)-1]
}

// Remove given argument form given slice, preserving order of the latter.
func ArgDelStable(slice []int, arg int) []int {
    return append(slice[:arg], slice[arg+1:]...)
}

// Benchmark comparison: emptying an 10^9—element array
// ArgDel:		 	 60 ps
// ArgDelStable:	224	 s
// https://stackoverflow.com/a/37335777