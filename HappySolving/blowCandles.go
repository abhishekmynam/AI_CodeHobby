package main

import "fmt"

func main(){
	var n,countMax int64
	fmt.Scan(&n)
	var candles = make ([]int64, n)
	for i,_ := range candles{
		fmt.Scan(&candles[i])
	}
	var max int64
	for _,j:=range candles{
		if(max <j){
			countMax=1
			max = j
		}else if (max == j){
			countMax++
		}
	}
	fmt.Println(countMax)
}
