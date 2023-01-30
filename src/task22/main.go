package main

import "fmt"

func main() {

	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	m := 0
	n := 3
	merge(nums1, m, nums2, n)
	fmt.Print(nums1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {

	if m == 0 {
		for i := 0; i < n; i++ {
			nums1[i] = nums2[i]
		}
		return
	}

	if n == 0 {
		return
	}

	indexNums1 := 0
	indexNums2 := 0

	var nums1Copy []int

	for i := 0; i < m; i++ {
		nums1Copy = append(nums1Copy, nums1[i])
	}

	for i := 0; i < len(nums1); i++ {
		val1 := nums1Copy[indexNums1]
		val2 := nums2[indexNums2]

		if val1 <= val2 {
			nums1[i] = val1
			indexNums1++
		} else {
			nums1[i] = val2
			indexNums2++
		}

		if indexNums1 == m {
			for j := i + 1; j < len(nums1); j++ {
				nums1[j] = nums2[indexNums2]
				indexNums2++
			}
			break
		}

		if indexNums2 == n {
			for j := i + 1; j < len(nums1); j++ {
				nums1[j] = nums1Copy[indexNums1]
				indexNums1++
			}
			break
		}

	}

}
