package main

import (
	"fmt"
)

func main(){
	var n,nextPos,lines int
	fmt.Scan(&n)
	var strs = make([]int, n)
	for j,_:= range strs{
		fmt.Scan(&strs[j])
	}

	prevPos := 0

	for i:=0 ; i<len(strs); i=nextPos{
		nextPos = prevPos +strs[i] + 1
		prevPos = nextPos
		lines++
	}
	fmt.Println(lines)
}
