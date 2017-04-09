package main

import (
	"fmt"

	"sort"

)

func main(){
	var lists= make ([]float64, 5)

	for k,_ := range lists{
		fmt.Scan(&lists[k])
	}

	sort.Float64s(lists)
	var min, max float64
	for i:=0; i<4;i++{
		min += lists[i]
	}
	minInt := int64(min)
	for j:=1;j<5; j++{
		max+= lists[j]
	}
	maxInt := int64(max)
	fmt.Print(minInt)
	fmt.Print(" ")
	fmt.Print(maxInt)
}
