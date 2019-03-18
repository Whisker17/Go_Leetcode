package main

import (
	"sort"
	"math"
)

func threeSumClosest(nums []int, target int) int {
	if len(nums) < 3 {
		return 0
	}
	var diff int
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if i == 0 {
			diff = nums[0] + nums[1] + nums[len(nums)-1] - target
		} else if nums[i-1] == nums[i] {
			continue
		}
		j, k := i+1, len(nums)-1
		for j < k {
			if j > i+1 && nums[j] == nums[j-1] {
				j++
				continue
			}
			tmp := nums[i] + nums[j] + nums[k] - target
			if tmp == 0 {
				return target
			} else {
				if tmp < 0 {
					j++
				} else {
					k--
				}
				if math.Abs(float64(tmp)) < math.Abs(float64(diff)) {
					diff = tmp
				}
			}
		}
	}
	return diff + target
}
