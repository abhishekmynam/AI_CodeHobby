package main

import "fmt"

func main(){
	var n,m int
	fmt.Scan(&n)
	var cntListN = make([]int, 10000)
	var cntListM = make([]int, 10000)
	var thisInt int
	//var nList = make ([] int, n)
	//var fList = make ([] int, 0)
	for i:=0;i<n;i++{
		fmt.Scan(&thisInt)
		cntListN[thisInt]++
	}
	fmt.Scan(&m)

	//var mList = make ([] int, m)
	for i:=0;i<m;i++{
		fmt.Scan(&thisInt)
		cntListM[thisInt]++
	}

	for p,q:= range cntListN{
		if (q<cntListM[p]){
			fmt.Print(p)
			fmt.Print(" ")
		}
	}

	/*for _,k := range fList{
		fmt.Print(k)
		fmt.Print(" ")
	}*/
}
