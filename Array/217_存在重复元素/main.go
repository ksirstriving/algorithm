package main

import "fmt"

//https://leetcode-cn.com/problems/contains-duplicate/
func main() {
    nums := []int{1, 2, 3, 1}
    solution(nums)
    fmt.Println(nums)
}

func solution(nums []int) bool {
    length := len(nums)
    if length <= 1 {
        return  false
    }

    maps := make(map[int]int)
    for i := 0; i < length; i++ {
        if _,ok := maps[nums[i]]; ok {
            return true
        } else {
            maps[nums[i]] = 1
        }
    }

    return  false
}