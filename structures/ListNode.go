package structures

import "fmt"

type ListNode struct {
    data interface{}
    next *ListNode
}

type LinkList struct {
    head *ListNode
    length uint
}

// 初始化链表
func NewLinkList() *LinkList {
    return &LinkList{
        head:   &ListNode{
            data: 0,
            next: nil,
        },
        length: 0,
    }
}

// 初始化节点数据
func NewLinkNode(value interface{}) *ListNode {
    return &ListNode{
        data: value,
        next: nil,
    }
}

// 获取下一个节点
func (n *ListNode) GetNext() *ListNode {
    return n.next
}

// 获取当前节点数据
func (n *ListNode) GetValue() interface{} {
    return n.data
}

// 尾部添加节点
func (l *LinkList) Append(node *ListNode) {
    cur := l.head
    for i := uint(0); i < l.length; i++ {
        cur = cur.next
    }
    cur.next = node
    l.length++
}

// 打印链表
func (l LinkList) Show() {
    cur := l.head.next
    for cur != nil {
        fmt.Print(cur.data)
        if cur.next != nil {
            fmt.Print(" -> ")
        }
        cur = cur.next
    }
}
//
//func main()  {
//    linkList := NewLinkList()
//    linkList.Append(NewLinkNode(1))
//    linkList.Append(NewLinkNode(2))
//    linkList.Append(NewLinkNode(3))
//    linkList.Show()
//}