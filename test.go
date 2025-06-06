package main

import "fmt"

var myMap = make(map[string]int)

func main(){
	myMap["age"]=1
	myMap["roll"]=12
	myMap["makrs"]=112

	for i,j := range myMap{
		fmt.Println(i,":",j)
	}
	// fmt.Println(myMap["age"])
}