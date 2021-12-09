// 在数组中找到 2 个数之和等于给定值的数字，结果返回 2 个数字在数组中的下标。
// Given nums = [2, 7, 11, 15], target = 9,
// Because nums[0] + nums[1] = 2 + 7 = 9,
// return [0, 1]
package main

import "fmt"

func main() {
	nums := []int{2, 1, 15, 7}
	target := 9
	result := twoSum(nums, target)
	fmt.Println(result)
}


func twoSum(num []int, target int) []int {
	m := make(map[int]int)
	fmt.Println(m)
	for i := 0; i < len(num); i++ {
		another := target - num[i]
		if _, ok := m[another]; ok {
			return []int{m[another], i}
		}
		m[num[i]] = i
		fmt.Println(m)
	}
	return nil
}
