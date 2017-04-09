package main

import (
	"fmt")

func main(){
	var n, t,  ncur, adds int

	fmt.Scan(&n, &t)
	cOut:= make([]int, t-1)
	ReadN(cOut, 0, t-1)
	ncur = n
	for i:=0 ; i<t ; i++{
		if(i+1 != t){
			ncur = ncur - cOut[i]
			if(ncur < 5){
				thisVal := n - ncur
				adds = adds+thisVal
				ncur = ncur + thisVal
			}
		}
	}
	fmt.Println(adds)
}
func ReadN(all []int, i, n int) {
	if n == 0 {
		return
	}
	if m, err := fmt.Scan(&all[i]); m != 1 {
		panic(err)
	}
	ReadN(all, i+1, n-1)
}
