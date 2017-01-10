package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	var matOrder int//,updown,leftright int
	var indexM,indexP, rowP,colP, rowM, colM  int
	var updown,leftright int
	var thisRow string
	fmt.Scan( &matOrder)
	room:= make([]string, matOrder)
	for i:=0 ;i<matOrder;i++{
		scanner.Scan()
		thisRow = scanner.Text()
		room[i]=thisRow
		indexM = strings.Index(thisRow, "m")
		indexP = strings.Index(thisRow, "p")
		if(indexM > -1){
			rowM = i
			colM = indexM
		}
		if(indexP > -1){
			rowP = i
			colP = indexP
		}

	}
	updown = rowP-rowM
	leftright = colP-colM
	if(updown >0 ){
		for i:=0;i<updown;i++{
			fmt.Println("DOWN")
		}
	}else if(updown <0){
		for i:=0;i<(0-updown);i++{
		fmt.Println("UP")
		}
	}
	if(leftright>0){
		for i:=0;i<leftright;i++{
			fmt.Println("RIGHT")
		}
	}else if (leftright<0){
		for i:=0;i<(0-leftright);i++{
			fmt.Println("LEFT")
		}
	}

}
