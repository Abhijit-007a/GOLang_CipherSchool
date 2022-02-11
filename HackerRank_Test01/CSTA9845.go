// Develop an encoding algorithm where consecutive data elements are replaced by just one data value and count.

// For example we can represent the original 53 characters with only 13.

// "WWWWWWWWWWWWBWWWWWWWWWWWWBBBWWWWWWWWWWWWWWWWWWWWWWWWB"  ->  "12WB12W3B24WB"
// RLE allows the original data to be perfectly reconstructed from the compressed data, which makes it a lossless data compression.

// "AABCCCDEEEE"  ->  "2AB3CD4E"  ->  "AABCCCDEEEE"
// Function naming format:-

// func RunLengthEncode(input string) string {} func RunLengthDecode(input string) string {}

// Input Format

// AABCCCDEEEE

// Constraints

// input is a string

// Output Format

// 2AB3CD4E

package main

import (
	"fmt"
	"strconv"
)

func RunLengthEncode(str string) string {
	n := len(str)
	storageString := ""
	for i := 0; i < n; i++ {
		var count int = 1
		for i < n-1 && str[i] == str[i+1] {
			count++
			i++
		}
		if count == 1 {
			storageString += string(str[i])
		} else {
			storageString += strconv.Itoa(count) + string(str[i])
		}
	}
	return storageString
}

func main() {
	var str string
	fmt.Scanln(&str)
	fmt.Println(RunLengthEncode(str))
}
