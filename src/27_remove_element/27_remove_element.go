package main

func removeElement(nums []int, val int) int {
	start := 0
	for i:= 0;i< len(nums);i++ {
		if nums[i] != val {
			start,nums[start] = start+1,nums[i]
		}
	}
	return start
}

