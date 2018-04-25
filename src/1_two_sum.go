package main

/*func twoSum(nums []int, target int) []int {
	l := len(nums)
	for i := 0;i<l;i++ {
		for j := i+1;j<l;j++ {
			if nums[i]+nums[j] == target {
				return []int{i,j}
			}
		}
	}

	return []int{}
}*/

func twoSum(nums []int, target int) []int {
	res := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		another := target - nums[i]
		if _, ok := res[another]; ok {
			return []int{res[another], i}
		}
		res[nums[i]] = i
	}
	return nil
}
