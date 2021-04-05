package _map

import (
	"fmt"
	"testing"
)

func lengthOfLongestSubstring(s string) int {
    if len(s) <= 0 {
    	return 0
	}

	sBytes := []byte(s)
	memo := [128]int{}
    maxLength := 0
    l := 0
    r := -1
    for l < len(sBytes) {
		if r + 1 < len(sBytes) && memo[sBytes[r + 1]] == 0 {
			r ++
			memo[sBytes[r]] ++
		} else {
			memo[sBytes[l]] --
			l ++
		}
		if r - l + 1 > maxLength {
			maxLength =  r - l + 1
		}
	}

	return maxLength
}

func Test(t *testing.T) {
    fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}
