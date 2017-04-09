package main

import (
	"fmt"
	"sort"
)

func main(){
	var m int
	result := "YES"
	fmt.Scan(&m)
	var s = make([]int,m)
	for i:=0; i<m; i++{
		var j int
		fmt.Scan(&j)
		s[i]=j
	}
	sort.Ints(s)
	for x:=1; x<m;x++{
		if(s[x]==s[x-1]){
			result = "NO"
			break
		}
	}
	fmt.Println(result)
}