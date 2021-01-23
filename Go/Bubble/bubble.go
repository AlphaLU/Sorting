package main

import (
	"fmt"
)

var p = fmt.Println

func main() {
	p("Started")
	list := []int{5, 1, 7, 9, 4, 2, 0, 8}
	arr := bubble(list)
	p(arr)
}

func bubble(arr []int) []int {
	var firstElement int = 0
	var secondElement int = 1
	for secondElement < len(arr) {
		if arr[firstElement] > arr[secondElement] {
			tmp := arr[firstElement]
			arr[firstElement] = arr[secondElement]
			arr[secondElement] = tmp
		}
		firstElement++
		secondElement++
	}
	out := check(arr)
	if out == false {
		bubble(arr)
	}
	return arr
}

func check(arr []int) bool {
	i := 0
	j := 1
	for j < len(arr) {
		if arr[i] > arr[j] {
			p("false")
			return false
		}
		i++
		j++
	}
	p("true")
	return true
}
