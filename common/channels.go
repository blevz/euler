package common

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func ReadLineByLine(path string, output chan<- string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		output <- scanner.Text()
	}
	close(output)
}

func Max_int(input <-chan int) int {
	max := int(math.MinInt32)
	for x := range input {
		if max < x {
			max = x
		}
	}
	return max
}

//accumulate

func Accumulate_int(input <-chan int) int {
	sum := 0
	for x := range input {
		sum += x
	}
	return sum
}

//Apply

func Apply_str(input <-chan string, output chan<- string, f func(string) string) {
	for x := range input {
		fmt.Println(x)
		output <- f(x)
	}
	close(output)
}

//Remove is opposite of apply
func Remove_int(input <-chan int, output chan<- int, f func(int) bool) {
	for x := range input {
		if !f(x) {
			output <- x
		}
	}
	close(output)
}
