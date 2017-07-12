package main

import "fmt"

func zero(z *int) {

	*z = 2
}

func main() {
	x := 5

	zero(&x)
	fmt.Println(x)
}
