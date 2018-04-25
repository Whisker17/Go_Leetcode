package main

import "math"

func maxArea(height []int) int {
	l := len(height)
	if l < 2 {
		return 0
	}
	maxArea1,left,right := 0,0,l-1
	for left < right {
		area := int(math.Min(float64(height[left]),float64(height[right]))) * (right-left)
		if area > maxArea1 {
			maxArea1 = area
		}
		if height[left] > height[right] {
			for right > 0 {
				right--
				if height[right] > height[right+1] {
					break
				}
			}
		} else {
			for left < l-1 {
				left++
				if height[left] > height[left-1] {
					break
				}
			}
		}
	}
	return maxArea1
}

