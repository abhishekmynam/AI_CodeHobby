package main

import (
	"fmt"
	"unicode/utf8"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)
	var stairs = make([]string, n)
	x := "#"
	var strs string
	for i := 1; i <= n; i++ {
		strs = strs + x
	}
	for i := 0; i < len(strs); i++ {
		y := strs[i:]
		z := strs[:i]
		w := strings.Repeat(" ", utf8.RuneCountInString(z))
		v := w + y
		stairs[n - i - 1] = v
	}
	for _, k := range stairs {
		fmt.Println(k)
	}
}
