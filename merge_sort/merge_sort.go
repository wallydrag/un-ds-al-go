// package merge_sort
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

	result := split(arr)
	fmt.Println("sorted array is : ", result)
}

func split(arr []int) []int {
	length := len(arr)
	if length == 1 {
		return arr
	}
	if length == 2 {
		sort(arr)
	}
	return merge(split(arr[0:length/2]), split(arr[length/2:length]))
}

func sort(arr []int) []int {
	if (len(arr) == 1) || (arr[0] <= arr[1]) {
		return arr
	}
	return []int{arr[1], arr[0]}
}

func merge(arr1, arr2 []int) []int {
	arr3 := make([]int, len(arr1)+len(arr2))
	for i := 0; i < len(arr3); i++ {
		if len(arr2) == 0 {
			copy(arr3[i:], arr1)
			break
		}
		if len(arr1) == 0 {
			copy(arr3[i:], arr2)
			break
		}
		if arr1[0] < arr2[0] {
			arr3[i] = arr1[0]
			arr1 = arr1[1:]
			continue
		}
		arr3[i] = arr2[0]
		arr2 = arr2[1:]
	}
	return arr3
}
