package main

import "fmt"

func main(){
	var t int
	fmt.Scan(&t)

	var trip = make([]trips,t)
	for i,_ := range trip{
		fmt.Scan(&trip[i].m)
		fmt.Scan(&trip[i].n)
		var thistrip = make ([]int,trip[i].n )
		for j,_ := range thistrip{
			fmt.Scan(&thistrip[j])
		}
		trip[i].costs = thistrip
	}

	for _,k := range trip{
		for p,q:= range k.costs{
			secVal := k.m - q
			secValIdx := contains(k.costs[p+1:],secVal)

			if(secValIdx!=-1){

				fmt.Print(p+1)
				fmt.Print(" ")
				fmt.Print(secValIdx+2+p)
				fmt.Println("")
				//fmt.Print(k.costs[p+1:])
				break
			}
		}
	}
}

type trips struct{
	m int
	n int
	costs []int
}

func contains(s []int, e int) int {
	for i, a := range s {
		if a == e {
			return i
		}
	}
	return -1
}