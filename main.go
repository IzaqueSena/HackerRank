package main

import "fmt"

func combinationsWithRepetions(n int32, a int32, b int32) [][]int32 {
	var out [][]int32
	var list []int32
	for i := int32(0); i < n; i++ {
		list = append(list, a)
	}
	for i := int32(0); i <= n; i++ {
		if i-1 >= 0 {
			list[i-1] = b
		}
		var copyList []int32
		copyOfList := make([]int32, len(list))
		copy(copyOfList, list)
		out = append(out, copyList)
	}
	return out
}

func main() {
	fmt.Println(combinationsWithRepetions(5, 0, 1))
}
