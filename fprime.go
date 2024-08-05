package ep

import (
	"strconv"
)

func IsPrime(nbr int) bool {
	if nbr <= 1 {
		return false
	}
	for i := 2; i*i <= nbr; i++ {
		if nbr%i == 0 {
			return false
		}
	}
	return true
}

func Fprime(nbr int) string {
	res := ""
	for i := 2; i <= nbr; i++ {
		if IsPrime(i) && nbr%i == 0 {
			res += strconv.Itoa(i) + "*"
			nbr /= i
			i = 2
		}
	}
	return res
}
