package common

import (
	"math"
	"strconv"
)

func Pow(n int, e int) int {
	product := 1
	for x := 0; x < e; x++ {
		product *= n
	}
	return product
}

func IsPalindrome(str string) bool {
	for x := 0; x < len(str)/2; x++ {
		if str[x] != str[len(str)-x-1] {
			return false
		}
	}
	return true
}

func IsPalindromeNum(num int) bool {
	str := strconv.Itoa(num)
	return IsPalindrome(str)
}

func Sqrt(num int) int {
	return int(math.Sqrt(float64(num)))
}

func IsPrime(num int) bool {
	upto := int(math.Sqrt(float64(num))) + 1
	for x := 2; x < upto; x++ {
		if num%x == 0 {
			return false
		}
	}
	return true
}

func IsOdd(num int) bool {
	return num%2 == 1
}

func IsEven(num int) bool {
	return num%2 == 0
}

func Max_int2(num1, num2 int) int {
	if num1 > num2 {
		return num1
	} else {
		return num2
	}
}

func genPrimesToNum(numMax int, output chan<- int) {
	for x := 2; x <= numMax; x++ {
		if IsPrime(x) {
			output <- x
		}
	}
	close(output)
}

func PrimeFactorization(num int) []int {
	factors := make([]int, 0)
	c := make(chan int)

	go genPrimesToNum(num, c)
	for x := range c {
		if num == 1 {
			break
		}

		for num%x == 0 {
			factors = append(factors, x)
			num = num / x
		}
	}
	return factors
}

func PrimeFactorizationMap(num int) map[int]int {
	factors := make(map[int]int)
	c := make(chan int)
	go genPrimesToNum(num, c)
	for x := range c {
		if num == 1 {
			break
		}

		for num%x == 0 {
			factors[x]++
			num = num / x
		}
	}

	return factors
}
