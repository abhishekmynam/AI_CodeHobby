package main

import (
	"fmt"
	"sort"
)

func main(){
	var n int
	var k float64
	fmt.Scan(&n,&k)

	var nList = make([]float64, n)
	for p,_ := range nList{
		fmt.Scan(&nList[p])
	}
	sort.Float64s(nList)
	i,j,cnt:=0,1,0

	for (j<n){
		diff := nList[j] - nList[i]
		if(diff == k){
			cnt++
			i++
			j++
		}else if(diff > k){
			i++
		}else if(diff < k ){
			j++
		}
	}
	fmt.Println(cnt)
}
