package main

import (
	"bufio"
	"os"
	"fmt"
	"reflect"
)
var(Room = make([][]string, 0,0)
	Visited =make([][]int,0,0)
	OutPut = make([][]int,0,0))
func main() {
	var matOrderX,matOrderY int
	var pacX, pacY, foodX, foodY int
	var thisRow string
	/*var Room = make([][]string, 0,0)
	var Visited =make([][]int,0,0)
	var OutPut = make([][]int,0,0)*/
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Scan(&pacX,&pacY)
	fmt.Scan(&foodX,&foodY)
	fmt.Scan(&matOrderX,&matOrderY)
	roomInput:= make([]string, matOrderX)
	/*Room := make([][]string, 0,0)
	Visited :=make([][]int,0,0)
	OutPut := make([][]int,0,0)*/
	coOdPoint:=make([]int,2)

	for i:=0; i<matOrderX; i++{
		scanner.Scan()
		thisRow = scanner.Text()
		roomInput[i]=thisRow
	}
	//fmt.Print(roomInput)
	for _,j:= range roomInput{
		tempRoom :=make([]string,0)
		for _,q:=range j{
			tempRoom = append(tempRoom,string(q))
		}
		Room = append(Room,tempRoom)
	}
	fmt.Print(Room)
	coOdPoint[1] = pacY
	coOdPoint[0] = pacX
	Visited = append(Visited,coOdPoint)

	for len(Visited)>0 {
		if (Room[Visited[0][0]][Visited[0][1]]!="."){
			pushPos(Visited[0],matOrderX,matOrderY)
			//curPos := make([]int,2)
			/*curPosX := Visited[0][0]
			curPosY := Visited[0][1]
			curPosX1 := curPosX+1
			curPosY1 := curPosY+1
			curPosX0 := curPosX-1
			curPosY0 := curPosY-1
			if curPosX0>-1 && (Room[curPosX0][curPosY]!="%"){
				coOdPoint[0] = curPosX0
				coOdPoint[1] = curPosY
				Visited = append(Visited, coOdPoint)
			}
			if curPosY0>-1 && (Room[curPosX][curPosY0]!="%"){
				coOdPoint[0] = curPosX
				coOdPoint[1] = curPosY0
				Visited = append(Visited, coOdPoint)
			}
			if curPosY1<matOrderY-1 && (Room[curPosX][curPosY1]!="%"){
				coOdPoint[0] = curPosX
				coOdPoint[1] = curPosY1
				Visited = append(Visited, coOdPoint)
			}
			if curPosX1<matOrderX-1 && (Room[curPosX1][curPosY]!="%"){
				coOdPoint[0] = curPosX1
				coOdPoint[1] = curPosY
				Visited = append(Visited, coOdPoint)
			}*/
			OutPut = append(OutPut,Visited[0])
			Visited = Visited[1:]
		}else if (Room[Visited[0][0]][Visited[0][1]]=="."){
			OutPut = append(OutPut,Visited[0])
			fmt.Println("Food Found")
			break
		}
	}
	fmt.Print(OutPut)
}

func pushPos(curPos []int,matOrderX int,matOrderY int){
	coOdPoint1:=make([]int,2)
	coOdPoint2:=make([]int,2)
	coOdPoint3:=make([]int,2)
	coOdPoint4:=make([]int,2)
	//if (Room[Visited[0][0]][Visited[0][1]]!="%"){
	if curPos[0]-1>-1 && (Room[curPos[0]-1][curPos[1]]!="%"){
		coOdPoint1[0] = curPos[0]-1
		coOdPoint1[1] = curPos[1]
		if(!ItemExists(coOdPoint1)) {
			Visited = append(Visited, coOdPoint1)
		}
	}
	if curPos[1]-1>-1 && (Room[curPos[0]][curPos[1]-1]!="%"){
		coOdPoint2[0] = curPos[0]
		coOdPoint2[1] = curPos[1]-1
		if(!ItemExists(coOdPoint2)) {
			Visited = append(Visited, coOdPoint2)
		}
	}
	if curPos[1]+1<matOrderY-1 && (Room[curPos[0]][curPos[1]+1]!="%"){
		coOdPoint3[0] = curPos[0]
		coOdPoint3[1] = curPos[1]+1
		if(!ItemExists(coOdPoint3)) {
			Visited = append(Visited, coOdPoint3)
		}
	}
	if curPos[0]+1<matOrderX-1 && (Room[curPos[0]+1][curPos[1]]!="%"){
		coOdPoint4[0] = curPos[0]+1
		coOdPoint4[1] = curPos[1]
		if(!ItemExists(coOdPoint4)) {
			Visited = append(Visited, coOdPoint4)
		}
	}

}
func ItemExists(curPos [] int)bool{
	for _,j := range OutPut{
		if(reflect.DeepEqual(j,curPos)){
			return true
		}
	}
	return false
}