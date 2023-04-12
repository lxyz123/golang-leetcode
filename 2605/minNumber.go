package main

import "sort"

func minNumber(nums1 []int, nums2 []int) int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	hash := make(map[int]int)
	for _, num := range nums1 {
		hash[num]++
	}
	for _, num := range nums2 {
		if _, ok := hash[num]; ok {
			return num
		}
	}
	a := nums1[0]
	b := nums2[0]
	if a < b {
		return a*10 + b
	} else {
		return b*10 + a
	}
}
