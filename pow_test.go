package main

// 2^10
// 4^5
// 4*(4^4)
// 4*(16^2)
// 4*(16*16)
func Pow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	if n < 0 {
		return 1 / Pow(x, n)
	}
	if n%2 == 0 {
		return Pow(x*x, n/2)
	} else {
		return Pow(x, n-1) * x
	}
}
