package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int)  {
	// No elements to be merged
	if m == 0 {
		copy(nums1, nums2)
		return
	} else if n == 0 {
		return
	}


	nums3 := make([]int, m + n)
	indexNum1 := 0;
	indexNum2 := 0;

	for i := 0; i < m + n; i++ {
		if indexNum2 >= n {
			nums3[i] = nums1[indexNum1];
			indexNum1 += 1;
			fmt.Println(1)
			continue
		}
		if nums1[indexNum1] == 0 {
			nums3[i] = nums2[indexNum2];
			indexNum2 += 1;
			indexNum1 += 1;
			fmt.Println(2)
			continue
		}

		if nums1[indexNum1] < nums2[indexNum2] {
			nums3[i] = nums1[indexNum1];
			indexNum1 += 1;
			fmt.Println(3)
		} else {
			nums3[i] = nums2[indexNum2];
			indexNum2 += 1;
			fmt.Println(4)
		}

	}
	copy(nums1, nums3)
}