package main

import (
    "fmt"
    "leetcode/structures"
)

var successor *structures.ListNode
func main() {
    // 1->2->3->4->5->6
    nums := []int{1,2,3,4,5,6}
    head := structures.Ints2List(nums)
    head.Show()
    //head = reverse(head)
    //head = reverseN(head, 3)
    head = reverseBetween(head, 2,4)
    n := structures.List2Ints(head)
    fmt.Printf("%v",n)
    //head.Show()
}


// 反转链表
func reverse(head *structures.ListNode) *structures.ListNode {
    if head == nil || head.Next == nil {
        return head
    }

    last := reverse(head.Next)
last.Show()
    head.Next.Next = head
    head.Next = nil
    return last
}

// 反转链表前n个节点
func reverseN(head *structures.ListNode, n int) *structures.ListNode {
    if n == 1 {
        // 记录第n+1节点
        successor = head.Next
        return head
    }
    // 以head.next为起点，反转前n-1节点
    last := reverseN(head.Next, n-1)
    head.Next.Next=head
    head.Next = successor
    return last

}

// 反正链表的一部分
func reverseBetween(head *structures.ListNode, m, n int) *structures.ListNode {
    if m == 1 {
        return reverseN(head, n)
    }

    head.Next = reverseBetween(head.Next, m-1, n-1)
    return head
}