package main

import "fmt"

func main() {
	n := greatest(1.5, 10, 72.2, 6, 5)
	fmt.Println(n)
}

func greatest(numbers ...float64) (float64){
	var max float64
	for _, v := range numbers {
		if v > max {
			max = v
		}
	}
	return max
}