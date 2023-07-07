package main

import (
	"fmt"
)

// 稳定
func bubbleSort(arr *[]int) {
	for i := 1; i < len(*arr); i++ {
		swap := false
		for j := 0; j < len(*arr)-1; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				(*arr)[j], (*arr)[j+1] = (*arr)[j+1], (*arr)[j]
				swap = true
			}
		}
		if !swap {
			return
		}
	}
}

// 稳定
func insertSort(arr *[]int) {
	for i := 1; i < len(*arr); i++ {
		t := (*arr)[i]
		j := i - 1
		for j >= 0 && (*arr)[j] > t {
			(*arr)[j+1] = (*arr)[j]
			j--
		}
		(*arr)[j+1] = t
	}
}

// 不稳定
func selectSort(arr *[]int) {
	for i := 0; i < len(*arr)-1; i++ {
		idx := i + 1
		for j := i + 1; j < len(*arr); j++ {
			if (*arr)[idx] > (*arr)[j] {
				idx = j
			}
		}
		(*arr)[i], (*arr)[idx] = (*arr)[idx], (*arr)[i]
	}
}

// 不稳定
func QuickSort(arr []int, start int, end int) {
	if start < end {
		mid := partition(arr, start, end)
		QuickSort(arr, start, mid-1)
		QuickSort(arr, mid+1, end)
	}
}

// 找到 pivot 位置
func partition(arr []int, start int, end int) int {
	pivot := arr[start]
	for start < end {
		for start < end && pivot <= arr[end] {
			end--
		}
		// 填补 low 位置空值
		arr[start] = arr[end]
		for start < end && pivot >= arr[start] {
			start++
		}
		// 填补end位置空值
		arr[end] = arr[start]
	}
	//pivot 填补 low位置的空值
	arr[start] = pivot
	return start
}

func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	i := len(arr) / 2
	left := mergeSort(arr[0:i])
	right := mergeSort(arr[i:])
	result := merge(left, right)
	return result
}

func merge(a1 []int, a2 []int) []int {
	i, j := 0, 0
	res := make([]int, 0)
	for i < len(a1) && j < len(a2) {
		if a1[i] < a2[j] {
			res = append(res, a1[i])
			i++
		} else {
			res = append(res, a2[j])
			j++
		}
	}
	for i < len(a1) {
		res = append(res, a1[i])
		i++
	}
	for j < len(a2) {
		res = append(res, a2[j])
		j++
	}
	return res
}

func main() {
	a := []int{5, 213, 5, 13, 5, 36, 2, 1, 4, 32, 42}
	res := mergeSort(a)
	fmt.Println(res)
}
