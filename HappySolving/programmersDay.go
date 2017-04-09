package main

import (
	"time"
	"fmt"
)
func main() {
	var curYear int
	fmt.Scan(&curYear)
	startDate := time.Date(curYear,time.January,01,00,00,00,00,time.UTC)
	progDate := startDate
	if(curYear<1918&&(curYear!=1700&&curYear!=1800&&curYear!=1900)){
		progDate= startDate.AddDate(0,0,255)
	}else if(curYear==1700||curYear==1800||curYear==1900){
		progDate= startDate.AddDate(0,0,254)
	}else if(curYear==1918){
		progDate= startDate.AddDate(0,0,255+13)
	}else if(curYear>1918){
		if((curYear%400==0)||(curYear%4==0&&curYear%100!=0)){
			progDate = time.Date(curYear,time.September,12,00,00,00,00,time.UTC)
		}else{
			progDate = time.Date(curYear,time.September,13,00,00,00,00,time.UTC)
		}
	}

	fmt.Println(progDate.Format("02.01.2006"))
}
