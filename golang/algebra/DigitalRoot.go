// Computes the digital root of a given integer.
func DigitalRoot(n int) int {
    return (n - 1) % 9 + 1
}