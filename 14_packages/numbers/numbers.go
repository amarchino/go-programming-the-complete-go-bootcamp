package numbers

import "math"

func Even(n uint) bool {
	return n%2 == 0
}

func odd(n int) bool {
	return n%2 != 0
}

func isPrime(num int) bool {
	sqrt := int(math.Pow(float64(num), 0.5))
	for i := 2; i < sqrt; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
func OddAndPrime(n int) bool {
	return odd(n) && isPrime(n)
}
