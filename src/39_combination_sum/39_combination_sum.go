package main

import "sort"

func combinationSum(candidates []int, target int) [][]int {
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
		ret = dfs(nums,target-num,i,sol,ret)
		sol = sol[:len(sol)-1]
	}
	return
}

/*
func combinationSum(candidates []int, target int) (results [][]int) {
	n := len(candidates)
	if n == 0 {
		return nil
	}

	for i := 0; i < n; i++ {
		if candidates[i] == target {
			results = append(results, []int{candidates[i]})
			continue
		}
		if candidates[i] > target {
			continue
		}

		res := combinationSum(candidates[i:], target-candidates[i])
		if res != nil {
			m := len(res)
			for j := 0; j < m; j++ {
				res[j] = append(res[j], candidates[i])
			}
			results = append(results, res...)
		}
	}
	return
}
*/
