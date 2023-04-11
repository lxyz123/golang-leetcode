package leetcode

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		if idx, ok := m[target-num]; ok {
			return []int{idx, i}
		}
		m[num] = i
	}
	return nil
}

func twoSum1(nums []int, target int) []int {
	for i, num := range nums {
		for j := i + 1; j < len(nums); j++ {
			if num+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}
