package main

import "fmt"

func main() {
	array := [6]int{3, 0, 0, 4, 0, 4}
	N := 6
	min := 0
	x := array[0]
	y := array[N-1]
	if x > y {
		min = y
	} else {
		min = x
	}
	sum := 0
	for i := 0; i < N; i++ {
		if array[i] <= min {
			sum = sum + (min - array[i])
		}
	}
	fmt.Println(sum)
}
