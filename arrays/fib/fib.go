package fib

func fib(N int) int {
	if N <= 1 {
		return N
	}
	i := 0
	a, b := 0, 1
	for i < N-1 {
		a, b = b, a+b
		i++
	}
	return b
}
