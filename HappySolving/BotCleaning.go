package main

import (
	"strings"
	"fmt"
	"reflect"
)

func GetPosOfAllCharsInString  (thisString string, tobesearched string) []int{
	var lastPos int
	positions := make([]int, 0)
	for strings.Index(thisString, tobesearched)>-1{
		lastPos= lastPos+strings.Index(thisString, "p")
		positions=append(positions, lastPos)
		thisString = thisString[strings.Index(thisString, "p")+1:]
		lastPos = lastPos+1
	}
	return positions
}

func ReturnOrderedPairAsc (orderSlice [][]int )[][]int{
	//var orderSlice = make([][] int, 4, 4)
	var iter,minrowSoFar int
	var sortedSlice = make([][] int, 0, 0)
	var nullify = []int{0}
	var minOrderSlice = make([] int, 2)

	/*var v1 = []int{1, 2}
	var v2 = []int{4,3}
	var v3 = []int{3, 4}
	var v4 = []int{4, 5}
	orderSlice[0] = v4
	orderSlice[1] = v3
	orderSlice[2] = v2
	orderSlice[3] = v1*/

	for(!AllSameStrings(orderSlice)){
		minrowSoFar=0
		var minSoFar,thisIter int
		for i,j:= range orderSlice{
			var thisSum int
			var baseSet bool =false
			for s,t:= range j {
				thisSum += t
				if !baseSet && iter==0 {
					minSoFar = thisSum
					if s!=0{baseSet = true}
				}else if !baseSet && iter!=0 && !reflect.DeepEqual(j,nullify){
					baseSet = true
				}else if !baseSet && iter!=0 && reflect.DeepEqual(j,nullify){
					baseSet = false
				}
			}
			if baseSet{
				if(thisIter == 0) {
					minSoFar = thisSum
					minrowSoFar = i
					minOrderSlice = j
				}else if(thisSum < minSoFar){
					minSoFar = thisSum
					minrowSoFar = i
					minOrderSlice = j
				}else if (thisSum == minSoFar && j[0] < minrowSoFar){
					minSoFar = thisSum
					minrowSoFar = i
					minOrderSlice = j
				} else if (thisSum == minSoFar && j[0] == minrowSoFar){
					minSoFar = thisSum
					minrowSoFar = i
					minOrderSlice = j
				}
				iter++
				thisIter++

			}
		}
		sortedSlice = append(sortedSlice, minOrderSlice)
		orderSlice[minrowSoFar] = nullify
	}

	return orderSlice
}

func printMoves (rowP int, rowM int, colP int, colM int){
	var updown, leftright int
	updown = rowP-rowM
	leftright = colP-colM
	if(updown >0 ){

		fmt.Println("DOWN")

	}else if(updown <0){
		fmt.Println("UP")

	}else if(updown ==0) {
		if (leftright > 0) {
			fmt.Println("RIGHT")
		}else if (leftright < 0) {
			fmt.Println("LEFT")
		}
	}
}