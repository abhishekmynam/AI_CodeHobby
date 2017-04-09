package main

import (
	"fmt"
	"strings"
	"reflect"
)
var SaveCoods = make([]Coords,0)
var VisitedCoods = make([]Coords,0)
var wandCnt int
func main(){
	var t int
	fmt.Scan(&t)
	var allPattern = make([]grid, t)
	for i,_:= range allPattern{
		fmt.Scan(&allPattern[i].m)
		fmt.Scan(&allPattern[i].n)
		for j:=0; j<allPattern[i].m ;j++{
			var thisVal string
			fmt.Scan(&thisVal)
			allPattern[i].gridVals = append(allPattern[i].gridVals, thisVal)
			starIdx := strings.Index(thisVal,"*")
			mIdx := strings.Index(thisVal,"M")

			if(starIdx!=-1){
				allPattern[i].posStrX =j
				allPattern[i].posStrY =starIdx
			}
			if(mIdx!=-1){
				allPattern[i].posMX = j
				allPattern[i].posMY = mIdx
			}
		}
		fmt.Scan(&allPattern[i].k)
	}

	for _,p := range allPattern{
		wandCnt = 0
		thisX, thisY:=p.posStrX,p.posStrY
		for (p.posMX!= thisX && p.posMY!=thisY){
			thisX,thisY = nextBestWay(thisX, thisY, p.gridVals)
			if(thisX == -99 && thisY == -99){
				length := len(SaveCoods)-1
				thisX =SaveCoods[length].x
				thisY =SaveCoods[length].y
				SaveCoods = SaveCoods[:length]
			}
		}
		if (wandCnt == p.k ){
			fmt.Println("Impressed")
		}else{
			fmt.Println("Oops!")
		}
	}



}

func nextBestWay(xCod int, yCod int, grid []string)(int, int){
	right := xCod + 1
	left := xCod - 1
	up := yCod + 1
	down := yCod - 1
	pathsHere := 0
	var thisCods Coords
	thisCods.x = xCod
	thisCods.y = yCod
	VisitedCoods = append(VisitedCoods,thisCods)
	thisCods.x = -99
	thisCods.y = -99
	var thisSaveCoods = make([]Coords,0)
	thisVal := reflect.TypeOf(grid[yCod][left])
	fmt.Println(thisVal)
	if(left>0){
		if(!didIVisit(left,yCod)) {
			//if (string(grid[yCod][left]) == ".") {
				thisCods.x = left
				thisCods.y = yCod
				thisSaveCoods = append(thisSaveCoods, thisCods)
				pathsHere++
			//}
		}
	}
	if(right<len(grid[yCod])){
		if(!didIVisit(right,yCod)) {
			//if (grid[yCod][right] == ".") {
				thisCods.x = right
				thisCods.y = yCod
				thisSaveCoods = append(thisSaveCoods, thisCods)
				pathsHere++
			//}
		}
	}
	if(up>0){
		if(!didIVisit(xCod,up)) {
			//if (grid[up][xCod] == ".") {
				thisCods.x = xCod
				thisCods.y = up
				thisSaveCoods = append(thisSaveCoods, thisCods)
				pathsHere++
			//}
		}
	}

	if(down<len(grid)){
		if(!didIVisit(xCod,down)) {
			//if (grid[down][xCod] == ".") {
				thisCods.x = xCod
				thisCods.y = down
				thisSaveCoods = append(thisSaveCoods, thisCods)
				pathsHere++
			//}
		}
	}

	if(pathsHere == 1){
		return thisCods.x, thisCods.y
	}else{
		for _,j := range thisSaveCoods{
			SaveCoods = append(SaveCoods,j)
		}
		wandCnt++
		return thisCods.x, thisCods.y
	}
}

func didIVisit(x int, y int)bool{
	for _,j:=range VisitedCoods{
		if(j.x == x && j.y == y){
			return true
		}
	}
	return false
}

type grid struct{
	m int
	n int
	gridVals []string
	k int
	posStrX int
	posStrY int
	posMX int
	posMY int
}

type Coords struct{
	x int
	y int
}
