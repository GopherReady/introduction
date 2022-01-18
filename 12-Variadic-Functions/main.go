package main

import (
	"fmt"
)

func sum(nums ...int) {
	total := 0
	for _, v := range nums {
		total += v
	}
	fmt.Println(total)
}

func main() {
	sum(1, 2)
	sum(1, 5, 6)
	nums := []int{1, 2, 3, 4, 5}
	sum(nums...)
}
