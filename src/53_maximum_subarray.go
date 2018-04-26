package main

func maxSubArray(nums []int) int {
	if len(nums) <= 0 {
		return -1
	}
	maxSum := nums[0]
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	for i := 1;i< len(nums);i++ {
		if dp[i-1] < 0 {
			dp[i] = nums[i]
		} else {
			dp[i] = dp[i-1] + nums[i]
		}
		if maxSum < dp[i] {
			maxSum = dp[i]
		}
	}
	return maxSum
}
