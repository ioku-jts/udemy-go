package main

import (
	"fmt"
)

func main() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, val := range arr {
		if val%2 == 1 {
			fmt.Println(val, "is Odd")
		} else {
			fmt.Println(val, "is Even")
		}
	}
}
