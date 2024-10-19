package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	sorted := sortSortedArrays(nums1, nums2)
	length := len(sorted)
	fmt.Printf("sorted: %v\n", sorted)
	if length%2 == 0 {
		return float64(sorted[length/2]+sorted[length/2-1]) / 2
	} else {
		return float64(sorted[length/2])
	}

}

func sortSortedArrays(nums1 []int, nums2 []int) []int {

	i, j := 0, 0
	result := make([]int, 0, len(nums1)+len(nums2))

	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			result = append(result, nums1[i])
			i++
		} else {
			result = append(result, nums2[j])
			j++
		}
	}

	// Append remaining elements of arr1
	for i < len(nums1) {
		result = append(result, nums1[i])
		i++
	}

	// Append remaining elements of arr2
	for j < len(nums2) {
		result = append(result, nums2[j])
		j++
	}

	return result
}
