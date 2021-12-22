package structures

import "fmt"

type ListNode struct {
    Data int
    Next *ListNode
}

type LinkList struct {
    head *ListNode
    length uint
}

// 链表转数组
func List2Ints(head *ListNode) []int {
    limit := 100
    times := 0
    res := []int{}

    for head != nil {
        times++
        if times > limit {
            msg := fmt.Sprintf("链表深度超过%d，可能出现环装链表。", limit)
            panic(msg)
        }
        res = append(res, head.Data)
        head = head.Next
    }
    return res
}

// 数组转链表
func Ints2List(nums []int) *ListNode {
    if len(nums) == 0 {
        return nil
    }

    l := &ListNode{}
    t := l
    for _, v := range nums{
        t.Next = &ListNode{
            Data: v,
        }
        t = t.Next
    }
    return l.Next
}

func (h *ListNode) Show() {
    if h == nil {
        return
    }
    fmt.Print(h.Data)
    for h.Next != nil {
        h = h.Next
        fmt.Print(" -> " , h.Data)
    }
    fmt.Println()
}