// package array_inversion
package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Enter size of your array: ")
	var size int
	fmt.Scanln(&size)
	var arr = make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Printf("Enter %dth element: ", i)
		fmt.Scanf("%d", &arr[i])
	}
	fmt.Println("Your intial array is: ", arr)

	_, result := sort_and_count_inversions(arr)
	fmt.Println("number of inversions: ", result)
}

func sort_and_count_inversions(arr []int) ([]int, int) {
	if len(arr) == 0 || len(arr) == 1 {
		return arr, 0
	}
	sorted_left, left_inversions := sort_and_count_inversions(arr[:len(arr)/2])
	sorted_right, right_inversions := sort_and_count_inversions(arr[len(arr)/2:])
	sorted_array, split_inversions := merge_and_count_split_inversions(sorted_left, sorted_right)
	return sorted_array, left_inversions + right_inversions + split_inversions
}

func merge_and_count_split_inversions(left, right []int) ([]int, int) {
	arr3 := make([]int, len(left)+len(right))
	counter := 0
	for i := 0; i < len(arr3); i++ {
		if len(right) == 0 {
			copy(arr3[i:], left)
			break
		}
		if len(left) == 0 {
			copy(arr3[i:], right)
			break
		}
		if left[0] < right[0] {
			arr3[i] = left[0]
			left = left[1:]
			continue
		}
		if left[0] == right[0] {
			arr3[i] = left[0]
			arr3[i+1] = right[0]
			left = left[1:]
			right = right[1:]
			i++
			continue
		}
		if left[0] > right[0] {
			arr3[i] = right[0]
			right = right[1:]
			counter = counter + len(left)
			continue
		}
	}
	return arr3, counter
}
