package main

import (
	"fmt"
)

func main() {
	_, b := multi()

	fmt.Println(b)

	arr := []int{5, 3, 8, 4, 2, 7, 5, 10}
	in := partition(arr, 0, len(arr))
	fmt.Println(in)
	fmt.Println(arr)
}

func partition(arr []int, l, h int) (pivotIndex int) {
	pivot := arr[l]
	i := l - 1
	j := h - 1
	for {

		for ok := true; ok; ok = !(arr[i] < pivot) {
			i++
		}

		for ok := true; ok; ok = !(arr[j] > pivot) {
			j--
		}
		if i >= j {

			pivotIndex = j
		}
		arr[i], arr[j] = arr[j], arr[i]
	}
	//return pivotIndex

}

func multi() (int, int) {
	return 5, 4
}
