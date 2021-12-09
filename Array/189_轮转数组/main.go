package main

import "fmt"

// https://leetcode-cn.com/problems/rotate-array/
func main() {
    //nums := []int{1, 2, 3, 4, 5, 6, 7}
    nums := []int{-1, -100, 3, 99}
    k := 2
    solution2(nums, k)
    fmt.Println(nums)
}

// 有问题
//func solution(nums []int, k int) {
//   length := len(nums)
//   k = k % length
//   if k == 0 {
//       return
//   }
//
//   tmp := make(map[int]int)
//   pos, stop := 0, length-k
//   for ; pos < length; pos++ {
//       if pos < stop {
//           // 从左往右移动 k
//           tmp[pos+k] = nums[pos+k]
//           if v, ok := tmp[pos]; ok {
//               nums[pos+k] = v
//           } else  {
//               nums[pos+k] = nums[pos]
//           }
//       } else {
//           //从右往左移动 length - k - 1
//           nums[pos-k-1], _ = tmp[pos]
//       }
//   }
//}

// 临时数组
func solution1(nums []int, k int) {
    length := len(nums)
    tmp := make(map[int]int)

    for i := 0; i < length; i++ {
        tmp[i] = nums[i]
    }
    for i := 0; i < length; i++ {
        p := (i + k) % length
        nums[p] = tmp[i]
    }
}

// 全部反转，然后先反转前k个，再反转后k个
func solution2(nums []int, k int) {
    length := len(nums)
    k %= length
    reverse(nums, 0, length-1)
    reverse(nums, 0, k-1)
    reverse(nums, k, length-1)
}

func reverse(nums []int, start int, end int) {
    for start < end {
        tmp := nums[start]
        nums[start] = nums[end]
        start++
        nums[end] = tmp
        end--
    }
}
