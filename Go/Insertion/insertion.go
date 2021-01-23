package main

func main() {
	a := []int{7, 6, 5, 4, 3, 2, 1}
	p(insertion(a))
}

func insertion(arr []int) []int {
	arrLen := len(arr)
	for i := 1; i < arrLen; i++ {
		key := arr[i]
		j := i - 1

		for j >= 0 && key < arr[j] {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}

	return arr
}