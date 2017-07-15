package main

import "fmt"


func main(){
	var smallnum int
	var largenum int

	fmt.Print("enter a large number: ")
	fmt.Scan(&largenum)
	fmt.Print("enter a small number: ")
	fmt.Scan(&smallnum)

	fmt.Println("the remainder of", largenum, "/" , smallnum, "is" , largenum % smallnum)



}