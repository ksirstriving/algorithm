package main

import "fmt"

//https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/
func main() {

    nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
    nLen := solution(nums)
    fmt.Printf("长度: %d %v\n", nLen, nums[:nLen])
}

func solution(nums []int) int {
    l := len(nums)
    if l < 2 {
        return l
    }
    length := 1
    left := 0
    for right := 1; right < l; right++ {
        if nums[left] != nums[right] {
            left++
            nums[left] = nums[right]
            length++
        }
    }
    return length
}
