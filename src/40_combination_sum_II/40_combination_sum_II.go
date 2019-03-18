package main

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	return dfs(candidates,target,0,[]int{},[][]int{})
}

func dfs(nums []int,target,start int,solution []int,res [][]int) (ret [][]int) {
	ret = res
	if target < 0 {
		return
	} else if target == 0 {
		ret = append(ret,solution)
		return
	}
	for i := start;i < len(nums);i++ {
		if i > start && nums[i] == nums[i-1] {
			continue
		}
		num := nums[i]
		sol := make([]int,len(solution))
		copy(sol,solution)
		sol = append(sol,num)
		ret = dfs(nums,target-num,i+1,sol,ret)
		sol = sol[:len(sol)-1]
	}
	return
}
