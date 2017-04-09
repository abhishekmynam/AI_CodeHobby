package main

import (
	"fmt"
	"strconv"
)

func main() {
	var sInt int
	fmt.Scan(&sInt)
	s:= strconv.Itoa(sInt)
	var count [9]int
	total := 0
	for n := 0; n < len(s); n++ {
		c := s[n]
		if(!('0' <= c && c <= '9')){
			panic("assertion failed: " )
		}
		var temp [9]int
		here := int(c - '0')
		for k, x := range count {
			temp[(k+here)%9] = x
		}
		if c != '0' {
			temp[here%9]++
		}
		count = temp
		if here % 2 == 0 {
			for k := 0; k < 9; k += 3 {
				total += count[k]
			}
		}
		if c == '0' {
			total++
		}
	}
	fmt.Println(total)
}