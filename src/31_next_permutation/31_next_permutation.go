package main

func nextPermutation(nums []int)  {
	l := len(nums)
	if l < 2 {
		return
	}
	index := l-1
	for index > 0 {
		if nums[index-1] < nums[index] {
			break
		}
		index--
	}
	if index == 0 {
		reverseNums(nums,0,l-1)
		return
	} else {
		val,j := nums[index-1],l-1
		for j >= index {
			if nums[j] > val {
				break
			}
			j--
		}
		nums[j],nums[index-1] = nums[index-1],nums[j]
		reverseNums(nums,index,l-1)
		return
	}
}

func reverseNums(nums []int,start,end int) {
	for i,j := start, end;i < j;i,j = i+1,j-1{
		nums[i],nums[j] = nums[j],nums[i]
	}
	return
}

