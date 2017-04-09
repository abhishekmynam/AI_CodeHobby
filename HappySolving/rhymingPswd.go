package main

import ("fmt")

var Vow = [...]string{"a","e","i","o","u"}
var Cons = [...]string{"b", "c", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p", "q", "r", "s", "t", "v", "x", "z", "w"}

func main() {
	var n int
	fmt.Scan(&n)

	pwdGen(n,true,"")
	pwdGen(n,false,"")
}

func pwdGen(n int, stat bool, str string){
	if n>0{
		if(stat){
			for _,cons:=range Cons{
				pwdGen(n-1,false, str+cons)
			}
		}else{
			for _,vows:=range Vow{
				pwdGen(n-1,true, str+vows)
			}
		}
	}else{
		fmt.Println(str)
	}
}
