package fib

// Fib fun
func Fib(n int) int {
	switch n {
	case 0, 1:
		return 1
	case 2:
		return 2
	default:
		return Fib(n-1)+Fib(n-2)
	}
}
