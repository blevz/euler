package main

import (
	"fmt"
	"github.com/blevz/euler/common"
	_ "math"
)

//Need to use principle of inclusion exclusion

type PFact struct {
	//Maps the prime factor with its exponent in the primefactorization
	data map[int]int
}

func (p PFact) Int() int {
	product := 1
	for k, v := range p.data {
		for x := 0; x < v; x++ {
			product *= k
		}
	}
	return product
}

func (p *PFact) SetFactorial(num, exp int) {
	p.data[num] = exp
}

func Make_PFact() *PFact {
	var p PFact
	p.data = make(map[int]int)
	return &p
}

func gcd(a, b int) int {
	for a != 0 {
		c := a
		a = b % a
		b = c
	}
	return b
}

func getResilentNumerators(denom int) int {
	m := make(map[int]struct{})
	primeFactorMap := common.PrimeFactorizationMap(denom)
	for k, _ := range primeFactorMap {
		for x := k; x < denom; x += k {
			m[x] = struct{}{}
		}
	}
	return denom - len(m) - 1
}

func getResilence(denom int) float64 {
	resilentNumerators := getResilentNumerators(denom)
	r := float64(resilentNumerators) / float64(denom-1)
	return r
}

func getResilentNumeratorsP(p *PFact) int {
	val := p.Int()
	num := val
	for k, _ := range p.data {
		num = num * (k - 1) / k
	}
	return num
}

func getResilenceP(p *PFact) float64 {
	resilentNumerators := getResilentNumeratorsP(p)
	r := float64(resilentNumerators) / float64(p.Int()-1)
	return r
}

func main() {
	fmt.Println(float64(15499) / float64(94744))
	//
	p := Make_PFact()
	p.SetFactorial(2, 3)
	p.SetFactorial(3, 1)
	p.SetFactorial(5, 1)
	p.SetFactorial(7, 1)
	p.SetFactorial(11, 1)
	p.SetFactorial(13, 1)
	p.SetFactorial(17, 1)
	p.SetFactorial(19, 1)
	p.SetFactorial(23, 1)

	fmt.Println(p.Int())
	r := getResilenceP(p)
	fmt.Println(r)
	if r < float64(15499)/float64(94744) {
		fmt.Println("fin", p.Int(), r)
	}

	//for x := 0; x < math.MaxInt32; x++ {
	/*minR := float64(1)
	for x := 2; x < math.MaxInt32; x++ {
		if !common.IsPrime(x) {
			p := Make_PFact()
			p.data = common.PrimeFactorizationMap(x)
			r := getResilenceP(p)
			if r < minR {
				minR = r
				fmt.Println(x)
			}

			if r < float64(15499)/float64(94744) {
				//if r < float64(4)/float64(10) {
				fmt.Println(x, r)

				break
			}
		}
	}*/
	/*
		fmt.Println(float64(15499) / float64(94744))
		c := make(chan int)
		go common.GenPrimes(c)
		num := 1
		minR := float64(1)
		for x := range c {
			fmt.Println(num)
			num *= x
			r := getResilence(num)
			if r < .2 {
				minR = r
				fmt.Println(num, r)
			}
		}
		fmt.Println(minR)*/
}
