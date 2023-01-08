package numbers

import (
	"math"
)

func IsPrime(num int) bool {
	sqrt := int(math.Pow(float64(num), 0.5))
	for i := 2; i < sqrt; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
