package main

import "fmt"

// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/
func main() {
    prices := []int{7, 1, 5, 3, 6, 4}
    //prices := []int{1,2,3}
    r := solution(prices)
    fmt.Println(r)
}

// 高抛低吸
func solution(prices []int) int {
    total := 0
    left := prices[0]

    for _, v := range prices[1:]{
        if v > left {
            total += v - left
        }
        left = v
    }

    return total
}
