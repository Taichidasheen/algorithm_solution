package main

import (
	"fmt"
)

/*
append操作：
he append built-in function appends elements to the end of a slice.
If it has sufficient capacity, the destination is resliced to accommodate（容纳） the new elements.
If it does not, a new underlying array will be allocated. Append returns the updated slice.
It is therefore necessary to store the result of append, often in the variable holding the slice itself
 */
func main() {
	arr := make([]int, 1, 5)
	s := append(arr, 2)
	fmt.Println(arr)
	fmt.Println(s)
	arr = append(arr, 3)
	fmt.Println(arr)
	fmt.Println(s)
}
