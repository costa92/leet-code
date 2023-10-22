package main

import "fmt"

func main() {
	nums := []int{0,0,1,1,1,2,2,3,3,4}

	result := removeDuplicatesV2(nums)
	//result := removeDuplicates(nums)
	fmt.Println(result)
}

func removeDuplicates(nums []int) int {
	lenNums := len(nums)
	if lenNums == 0 {
		return 0
	}
	left, right := 1, 1
	for right < lenNums {
		if nums[right] != nums[right-1] {
			nums[left] = nums[right]
			left++
		}
		right++
	}
	return left
}

func removeDuplicatesV2(nums []int) int {
	lenNums := len(nums)
	if lenNums == 0 {
		return 0
	}
	slow := 1
	for i := 0; i < lenNums - 1; i++ {
		if nums[i] != nums[i+1] {
			nums[slow] = nums[i+1]
			slow++
		}
	}
	return slow
}
