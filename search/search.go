package main

import "fmt"

// 二分查找
func bsearch(arr []int, n int) int {
	l, r := 0, len(arr)-1
	for l < r {
		mid := (l + r) / 2
		if arr[mid] == n {
			return mid
		} else if arr[mid] < n {
			mid = l
		} else {
			r--
		}
	}
	return -1
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(bsearch(a, 5))
}
