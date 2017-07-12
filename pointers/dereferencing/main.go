package main

import "fmt"

func main() {

	a := 50

	fmt.Println(a)

	var b = &a

	fmt.Println(b)
	fmt.Println(*b)

	*b = 20

	fmt.Println(a)

}
