package main

import "fmt"

func main(){
	var m,n int
	fmt.Scan(&m,&n)
	var swaps = make([]swapItem,n)
	for i,_ := range swaps{
		fmt.Scan(&swaps[i].x,&swaps[i].y)
	}

	for _,j := range swaps{
		if(j.x == m){
			m = j.y
		}else if(j.y==m){
			m = j.x
		}
	}
	fmt.Println(m)
}
type swapItem struct{
	x int
	y int
}
