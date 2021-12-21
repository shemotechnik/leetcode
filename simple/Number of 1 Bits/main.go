package main

import (
	"fmt"
)

func main() {
	fmt.Println(hammingWeight(100))
	fmt.Println(hammingWeight(253))
	fmt.Println(hammingWeight(2342342))
}

func hammingWeight(num uint32) int {
	str := fmt.Sprintf("%b", num)
	l := 0
	r := len(str) - 1
	cnt := 0
	for l <= r {
		if string(str[l]) == "1" {
			cnt++
			l++
		} else if string(str[r]) == "1" {
			cnt++
			r--
		} else {
			l++
			r--
		}
	}
	return cnt
}
