package main

import "fmt"

func main() {
	fmt.Println(greet("Jane", " doe"))
}

func greet(fname, lname string) string {
	return fmt.Sprint(fname, lname)
}
