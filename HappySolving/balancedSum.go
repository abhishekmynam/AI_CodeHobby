package main

import (
	"fmt"
	"math"
)

func main(){
	var t int
	fmt.Scan(&t)
	var allLists = make([]lists,t)

	for i,_:=range allLists{
		fmt.Scan(&allLists[i].itemsNum)
		var thisList = make([]int, allLists[i].itemsNum)
		for j,_:=range thisList{
			fmt.Scan(&thisList[j])
		}
		allLists[i].items = thisList
	}

	for _,p := range allLists{
		var nextIdx int
		var pilot, prevPilot int
		startIdx := int(math.Floor(float64(p.itemsNum-1)/2))
		for i:=startIdx;i<p.itemsNum;i=nextIdx{
			leftSlice:= balancedSum(p.items[:i])
			rightSlice:= balancedSum(p.items[i+1:])
			if(leftSlice>rightSlice){
				nextIdx = i-1
				pilot = -1
			}else if(leftSlice<rightSlice){
				nextIdx = i+1
				pilot = 1
			}else if(leftSlice == rightSlice){
				fmt.Println("YES")
				break
			}
			if(prevPilot != pilot && i!=startIdx){
				fmt.Println("NO")
				break
			}
			prevPilot = pilot
		}
	}


}

func balancedSum(thisList []int)(int){
	var thisSum int
	for _,j:=range thisList{
		thisSum += j
	}
	return thisSum
}

type lists struct{
	itemsNum int
	items []int
}

