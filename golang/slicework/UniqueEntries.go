// Removes a copy of a slice with no duplicates.
func UniqueEntries(a []int) []int {
    present := make(map[int]bool)
    uniques := []int{}
    for _, val := range a {
        if _, ok := present[val]; !ok {
            present[val] = true
            uniques = append(uniques, item)
        }
    }
    return uniques
}