package main

import (
	"fmt"
	"sort"
)

func main() {
	var n,k,txrCnt int
	fmt.Scan(&n, &k)
	var houses = make([]int, n)
	//var towers = make([]int, 0)

	for i,_ := range houses{
		fmt.Scan(&houses[i])
	}

	sort.Ints(houses)
	var minVal int
	for _,j:= range houses{
		if(j>minVal){
			txrCnt++
			minVal = j+2*k

		}
	}
	fmt.Print(txrCnt)
}


