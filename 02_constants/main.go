package main

import "fmt"

const (
	_  = iota
	KB = (iota * 10)
	MB = (iota * 10)
	GB = (iota * 10)
)

func main() {
	fmt.Println("binary\t\t\tdecimal")
	fmt.Printf("%b\t", KB)
	fmt.Printf("%d\n", KB)
	fmt.Printf("%b\t", MB)
	fmt.Printf("%d\n", MB)
	fmt.Printf("%b\t", GB)
	fmt.Printf("%d\n", GB)

}
