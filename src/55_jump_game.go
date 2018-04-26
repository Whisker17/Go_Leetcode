package main

func canJump(nums []int) bool {
	l := len(nums)
	maxRight := 0
	for i := 0;i < l;i++ {
		if i <= maxRight {
			maxRight = int(math.Max(float64(maxRight),float64(i+nums[i])))
		}
		if maxRight >= len(nums)-1 {
			return true
		}
	}
	return false
}
