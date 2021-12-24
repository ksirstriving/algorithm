package main

import "leetcode/structures"

func main() {
    // 1->2->3->4->5->6
    nums := []int{1,2,3,4,5,6}
    head := structures.Ints2List(nums)
    head.Show()
    //head = reverse(head)
    //head = reverseBetween(head,head.Next.Next.Next)
    head = reverseKGroup(head, 4)
    head.Show()
}

// 迭代方式实现反转链表
func reverse(head *structures.ListNode) *structures.ListNode {
    var pre *structures.ListNode
    cur := head
    nxt := head
    for cur != nil {
        nxt = cur.Next
        // 反转节点
        cur.Next = pre
        // 更新指针
        pre = cur
        cur = nxt
    }
    // 返回反转后头结点
    return pre
}

// 迭代方式实现反转链表 区间[a,b)
func reverseBetween(a, b *structures.ListNode) *structures.ListNode {
    var pre *structures.ListNode
    cur := a
    nxt := a
    for cur != b {
        nxt = cur.Next
        // 反转节点
        cur.Next = pre
        // 更新指针
        pre = cur
        cur = nxt
    }
    // 返回反转后头结点
    return pre
}

// k个一组反转链表
func reverseKGroup(head *structures.ListNode, k int) *structures.ListNode {
    if head == nil {
        return head
    }
    //区间[a,b)包含k个待反转节点
    a := head
    b := head
    for i := 0; i < k; i++ {
        if b == nil {
            // 不足k个不反转
            return head
        }
        b = b.Next
    }
    // 反转前k个
    newHead := reverseBetween(a,b)
    // 递归
    a.Next = reverseKGroup(b, k)
    return newHead
}