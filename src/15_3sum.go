package main

import "sort"

func threeSum(nums []int) [][]int {
	var res [][]int
	if len(nums) < 3 {
		return res
	}
	sort.Ints(nums)
	for i := 0;i<len(nums)-2;i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j,k := i+1,len(nums)-1
		for j < k {
			if j > i+1 && nums[j] == nums[j-1] {
				j++
				continue
			}
			if nums[i] + nums[j] + nums[k] == 0 {
				res = append(res,[]int{nums[i],nums[j],nums[k]})
				j++
				k--
			} else if nums[i] + nums[j] + nums[k] > 0 {
				k--
			} else {
				j++
			}
		}
	}
	return res
}
