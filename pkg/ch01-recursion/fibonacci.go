package ch01

func Fibo(n int) int {
	if n == 0 {
		return 0
	} else if n < 3 {
		return 1
	} else {
		return Fibo(n-1) + Fibo(n-2)
	}
}
