package main

import (
	"fmt"
)

var p = fmt.Println

func main() {
	array := []int{5, 4, 2, 7, 8, 1, 3, 6}
	newarr := mergeSort(array)
	p(newarr)
}

func merge() {

}

func mergeSort(arr []int) []int {
	if len(arr) > 1 {
		mid := (len(arr)) / 2

		L := arr[:mid]
		R := arr[mid:]

		mergeSort(L)
		mergeSort(R)

		i, j, k := 0, 0, 0

		for i < len(L) && j < len(R) {
			if L[i] < R[j] {
				arr[k] = L[i]
				i++
			} else {
				arr[k] = R[j]
				j++
			}
			k++
		}

		for i < len(L) {
			arr[k] = L[i]
			i++
			k++
		}

		for j < len(R) {
			arr[k] = R[j]
			j++
			k++
		}
	}

	return arr

}
