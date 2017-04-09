package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	var retStr string
	fmt.Scan(&n)
	funcStr := "min(int, int)"
	retStr = buildStr(n,funcStr)
	fmt.Println(retStr)
}

func buildStr(n int, funcStr string)string{
	var thisStr string
	if(n>2){
		thisStr = buildStr(n-1, funcStr)
	}
	i := strings.LastIndex(thisStr,"int")
	if(i==-1){
		thisStr = funcStr
	}else {
		thisStr = thisStr[:i] + funcStr + thisStr[i + 3:]
	}
	return thisStr
}

