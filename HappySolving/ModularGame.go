package main

import "fmt"

func main(){
	var n,a,b, sums, minVal, prevCnt, minVal1 int
	fmt.Scan(&n,&a,&b)
	var alice = make([]int, a)
	var bob = make([]int, b)
	var sumVals = make([]int, a*b)
	//var cnts = make([]int, n)

	for i, _ := range alice {
		fmt.Scanf("%d", &alice[i])
	}
	for j, _ := range bob {
		fmt.Scanf("%d", &bob[j])
	}
	for _,k :=range alice{
		for _,l:=range bob{
			sumVals[sums]=(k+l)
			sums++
		}
	}
	for m:=1; m<=n; m++{
		minVal = 0
		thisVal := m
		cnt := 0
		for _,p := range sumVals{
			if((p+m)%n==0){
				cnt++
				minVal = -1
				if(m==1){
					prevCnt = cnt
					minVal1 = thisVal
				}else if(cnt >= prevCnt){
					thisVal = minVal1
					break
				}
			}
		}
		prevCnt = cnt
		minVal1 = thisVal

		if(minVal == 0){
			minVal = m
			break
		}
	}
	/*if(minVal == -1){
		mins :=cnts[0]
		ns := 0
		for w,y:=range cnts{
			if(y<mins){
				mins = y
				ns = w
			}
		}
		minVal = ns+1
	}*/

	fmt.Println(minVal)
}
