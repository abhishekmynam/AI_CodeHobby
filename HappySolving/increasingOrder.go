package main

import (
	"fmt"
	"sort"
)

type strInts struct {
	vals string
	lens int
}

type ByLength []string

func (s ByLength) Len() int {
	return len(s)
}
func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ByLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}
func main() {
	var  n int
	fmt.Scan(&n)
	var ints = make ([]string, n)
	for i,_ := range ints{
		fmt.Scan(&ints[i])
	}
	sort.Sort(ByLength(ints))

	var prev, index int
	prev = len(ints[0])
	for j,k := range ints{
		if(len(k)> prev){
			thisSlice := ints[index:j]
			sort.Strings(thisSlice)
			for _,p:=range thisSlice{
				fmt.Println(p)
			}
			prev = len(k)
			index = j
		}
		if(j == len(ints)-1){
			thisSlice := ints[index:]
			sort.Strings(thisSlice)
			for _,p:=range thisSlice{
				fmt.Println(p)
			}
		}
	}
}
