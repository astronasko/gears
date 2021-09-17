// Computes the GCD of two numbers.
func LeastCommonMultiple(a,b int) int {
    return a*b/GCD(a,b)
}