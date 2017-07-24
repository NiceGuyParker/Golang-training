package main

import "fmt"

func small() int {
	smallest := 0
	for i := 0; i <= 1000000; i++ {
		if i%1 == 0 {
			if smallest < i {
				smallest = i
			}
		}
		if i%2 == 0 {
			if smallest < i {
				smallest = i
			}
		}
		if i%3 == 0 {
			if smallest < i {
				smallest = i
			}
		}
		if i%4 == 0 {
			if smallest < i {
				smallest = i
			}
		}
		if i%5 == 0 {
			if smallest < i {
				smallest = i
			}
		}
		if i%6 == 0 {
			if smallest < i {
				smallest = i
			}
		}
		if i%7 == 0 {
			if smallest < i {
				smallest = i
			}
		}
		if i%8 == 0 {
			if smallest < i {
				smallest = i
			}
		}
		if i%9 == 0 {
			if smallest < i {
				smallest = i
			}
		}
		if i%10 == 0 {
			if smallest < i {
				smallest = i
			}
		}

	}
	return smallest
}

func main() {
	n := small

	fmt.Printf("%d", n)
}
