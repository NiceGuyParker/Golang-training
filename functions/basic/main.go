package main

import "fmt"

func main() {
	greet("Dierdre", " Smith!")

}
func greet(fname, lname string) {
	fmt.Print(fname,lname)
}
