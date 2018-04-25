package main

import "sort"

func fourSum(nums []int, target int) [][]int {
	if len(nums) < 4 {
		return 0
	}
	sort.Ints(nums)
	var res [][]int
	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums)-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			left, right := j+1, len(nums)-1
			for left < right {
				if left > j+1 && nums[left] == nums[left-1] {
					left++
					continue
				}
				if nums[i]+nums[j]+nums[left]+nums[right] == target {
					res = append(res, []int{nums[i], nums[j], nums[left], nums[right]})
					left++
					right--
				} else if nums[i]+nums[j]+nums[left]+nums[right] < target {
					left++
				} else {
					right--
				}
			}
		}
	}
	return res
}
