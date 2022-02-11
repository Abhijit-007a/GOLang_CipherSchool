package main

import "fmt"

func Fibonacci() {
	fmt.Println(fibo(4))
}
func fibo(number int) int {
	if number == 0 || number == 1 {
		return number
	}
	result := fibo(number-1) + fibo(number-2)
	return result
}
