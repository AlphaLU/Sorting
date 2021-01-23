package main

import (
	"sort"
	"testing"
)

func TestMergeSort(t *testing.T) {
	cases := [][]int{{1,2,3,4,5,6,7,8,9},{9,8,7,6,5,4,3,2,1,0},{139, 320, 204, 186, 309, 261, 189, 90, 223, 322},{174, 44, 79, 243, 55, 162, 217, 220, 6, 84},{-19, -8, -18, -76, -12, -42, -20, -33, -84, 0},{-73, -66, -63, -27, -43, -92, -74, -90, -26, -22}}
	for i := 0; i < len(cases); i++ {
		output := mergeSort(cases[i])
		
		if !sort.IntsAreSorted(output) {
			t.Errorf("Input: %v\n", cases[i])
			t.Errorf("Output: %v\n", output)
		} else {
			p("Unsorted: ",cases[i])
			p("Sorted: ", output)
		}
	}
}