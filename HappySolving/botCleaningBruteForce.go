package main
import (
	"fmt"
	"bufio"
	"os"
)
func main() {
	var matOrder int =5
	var colM, rowM int
	var thisRow string
	scanner := bufio.NewScanner(os.Stdin)
	room:= make([]string, matOrder)
	fmt.Scan(&rowM,&colM)

	for i:=0; i<matOrder; i++{
		scanner.Scan()
		thisRow = scanner.Text()
		room[i]=thisRow
	}
	printFn(room,matOrder)


}


func printFn (room []string, matOrder int){
	var dir int =1
	for i,j := range room{
		if(dir == -1){
			j = Reverse(j)
		}
		for _,q := range j{
			if(dir==1){if(q=='d'){
				fmt.Println("CLEAN")
			}else if(q=='-'){
				fmt.Println("RIGHT")}
			}
			if(dir==-1){if(q=='d'){
				fmt.Println("CLEAN")
			}else if(q=='-'){
				fmt.Println("LEFT")}
			}
		}
		dir = dir*-1

		if(i<matOrder-1){fmt.Println("DOWN")}
	}
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}