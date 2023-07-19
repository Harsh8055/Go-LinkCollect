package main

import (
	"fmt"
)

func main() {

fmt.Printf("Getting Started, initializing")

arr := []int{9, 7, 3, 5,3, 2}

mergeSort:= mergeSort(arr);
fmt.Printf("Sorted Array, %v", mergeSort)

}

// merge two sorted array, returns sorted array
func merge(arr1 []int, arr2 []int) []int {
	var mergedArr []int;
	pointer1 := 0
	pointer2 := 0

	k := 0


	for pointer1 < len(arr1) && pointer2 < len(arr2) {
		if arr1[pointer1] < arr2[pointer2] {
			mergedArr = append(mergedArr, arr1[pointer1])
			pointer1++
		} else if arr1[pointer1] > arr2[pointer2] {
			mergedArr = append(mergedArr, arr2[pointer2])
			pointer2++
		} else {
			mergedArr = append(mergedArr, arr1[pointer1])
			mergedArr = append(mergedArr, arr1[pointer1])
			pointer1++
			pointer2++
		}
		k++
	}

	for pointer1 < len(arr1) {
		mergedArr = append(mergedArr, arr1[pointer1])
		pointer1++
		k++;
	}

	for pointer2 < len(arr2) {
		mergedArr = append(mergedArr, arr2[pointer2])
		pointer2++
		k++;
	}
	

	fmt.Printf("mergedArr %v", mergedArr)

	return mergedArr

}


func mergeSort(arr []int) []int {

	if (len(arr) < 2) {
		return arr
	}
	mid := len(arr) / 2
	fmt.Printf("mid %v \n",mid)
    
	arr1 := arr[:mid]
	arr2 := arr[mid:]
	fmt.Printf("arr %v %v \n",arr1, arr2)
	return merge(mergeSort(arr1), mergeSort(arr2));

}