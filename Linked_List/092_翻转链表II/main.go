package main

type Object interface {}

type Node struct {
    Data Object
    Next *Node
}

type List struct {
    headNode *Node
}