package main

import (
	"fmt"
)

func main() {
	arr := make([]int, 1, 5)
	s := append(arr, 2)
	fmt.Println(arr)
	fmt.Println(s)
	arr = append(arr, 3)
	fmt.Println(arr)
	fmt.Println(s)
}
