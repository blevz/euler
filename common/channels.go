package common

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

//ReadLineByLine opens the file specificied by path and puts each
//Line encountered into output
func ReadLineByLine(path string, output chan<- string) {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		output <- scanner.Text()
	}

	close(output)
}

//Max_int returns the largest int passed into the input channel
func Max_int(input <-chan int) int {
	max := int(math.MinInt32)
	for x := range input {
		if max < x {
			max = x
		}
	}
	return max
}

//Accumulate_int returns the sum of everything passed on the
//provided channel input once it has been closed
func Accumulate_int(input <-chan int) int {
	sum := 0
	for x := range input {
		sum += x
	}
	return sum
}

//Apply_str takes an input chan and outputs the result
func Apply_str(input <-chan string, output chan<- string, f func(string) string) {
	for x := range input {
		fmt.Println(x)
		output <- f(x)
	}
	close(output)
}

//Remove_int passes ints from input to output if they don't return true
//when passed to the supplied function f
func Remove_int(input <-chan int, output chan<- int, f func(int) bool) {
	for x := range input {
		if !f(x) {
			output <- x
		}
	}
	close(output)
}
