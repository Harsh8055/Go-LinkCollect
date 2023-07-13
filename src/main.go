package main

import (
	"fmt"
)

func main() {

fmt.Printf("Getting Started, initializing")

arr := []int{9, 7, 3, 5, 2}

mergeSort:= mergeSort(arr);
fmt.Printf("Sorted Array, %-1t", mergeSort)

}

// merge two sorted array, returns sorted array
func merge(arr1 []int, arr2 []int) []int {
	var mergedArr []int;
	pointer1 := 0
	pointer2 := 0
	length := 0

	if(len(arr1) > len(arr2)) {
		length = len(arr2)
	} else if(len(arr1) < len(arr2)){
		length = len(arr1)
	} else { length = len(arr1)}

	for i := 0; i < length; i++ {
		
		if(arr1[pointer1] < arr2[pointer2]) {
			mergedArr = append(mergedArr, arr1[pointer1])
		
			pointer1++
		} else if(arr1[pointer1] > arr2[pointer2]) {
			mergedArr = append(mergedArr, arr2[pointer2])
			pointer2++
		} else {
			mergedArr = append(mergedArr, arr1[pointer1])
			mergedArr = append(mergedArr, arr1[pointer1])
			pointer1++
			pointer2++
		}

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