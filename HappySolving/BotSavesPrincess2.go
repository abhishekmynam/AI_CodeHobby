
package main

import (
"fmt"
"bufio"
"os"
"strings"
)

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	var matOrder int
	var indexM,indexP, rowP,colP, rowM  int
	var updown,leftright int
	var thisRow string
	fmt.Scan( &matOrder)
	fmt.Scan(&rowM,&indexM)
	room:= make([]string, matOrder)
	for i:=0 ;i<matOrder;i++{
		scanner.Scan()
		thisRow = scanner.Text()
		room[i]=thisRow
		indexP = strings.Index(thisRow, "p")
		if(indexP > -1){
			rowP = i
			colP = indexP
		}
	}
	updown = rowP-rowM
	leftright = colP-indexM
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
