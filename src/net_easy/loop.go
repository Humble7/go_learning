package main

import (
	"fmt"
	"strconv"
)

func convertToBin(n int) string {
	result := ""
	for ; n > 0; n >>= 1 {
		result = strconv.Itoa(n & 1) + result
	}

	return result
}

func main() {
	fmt.Println(
		convertToBin(5),  // 101
		convertToBin(13), // 1101
		)
}
