package ch6alg_test

import (
	"fmt"
	"testing"
)

// LinkNode is a node in a doubly linked list
// after usage, LinkNode must be released like linkNode = nil
type LinkNode struct{
	Val int
	Next *LinkNode
	Pre *LinkNode
}

// Insert to tail
func (node *LinkNode) InsertTail(link *LinkNode,val int) {
	for link.Next !=nil{
		link = link.Next
	}
	link.Next = &LinkNode{Val: val}
}

// Insert to head
func (node *LinkNode) InsertHead(link *LinkNode, val int) *LinkNode {
     head := &LinkNode{Val: val, Next: link}
	 return head
}


// Delete a Node
func (node *LinkNode) DeleteNode(link *LinkNode, val int) *LinkNode {
	if link == nil{
		return nil
	}
	head := link
	pre := link
	for link != nil {
	    if link.Val == val{
			if link == head { // if the first node is to be deleted
			    head = link.Next
				link = nil
				return head
			}
			pre.Next = link.Next // delete the node
			link = nil // release memory
			return head
		}
		pre = link
		link = link.Next
	}
	return head
}

// Reverse the link
func (node *LinkNode) Reverse(link *LinkNode) *LinkNode{ 
	if link == nil {
		return nil
	}
	pre := link
	back := link.Next
	reveLink := link.Next
	pTmp := link

	for link != nil{
		if back == nil{
			return reveLink
		}
		
		link = back.Next
		back.Next = pTmp // reverse the link
		reveLink = back
		pTmp = back
		back = link
	}
	pre.Next = nil
	return reveLink
}

// 链表中环的检测
func (node *LinkNode) HasCycle(link *LinkNode) bool {
	if link == nil || link.Next == nil{
		return false
	}
	slow := link
	fast := link.Next

	for fast != nil && fast.Next != nil{
		if slow == fast {
			return true
		}

		slow = slow.Next
		fast = fast.Next.Next
	}
	return false
}


// 合并两个有序链表
func MergeTwoLists(l1 *LinkNode, l2 *LinkNode) *LinkNode {
	if l1 == nil && l2 != nil{
		return l2
	}
	if l1 != nil && l2 == nil{
		return l1
	}

	if l1.Val < l2.Val{
		l1.Next = MergeTwoLists(l1.Next,l2)
		return l1
	}else{
		l2.Next = MergeTwoLists(l1,l2.Next)
		return l2
	}

}


func TestLinkInsertTail(t *testing.T) {
	linkNode := &LinkNode{Val: 1}
	linkNode.InsertTail(linkNode, 2)
	linkNode.InsertTail(linkNode, 3)
	linkNode.InsertTail(linkNode, 4)
	// linkNode=linkNode.DeleteNode(linkNode,4)
	// for linkNode != nil {
	// 	fmt.Printf("%d ", linkNode.Val)
	// 	linkNode = linkNode.Next
	// }
	linkNode2 := &LinkNode{Val: 4}
	linkNode2.InsertTail(linkNode2, 5)
	linkNode2.InsertTail(linkNode2, 6)
	linkNode2.InsertTail(linkNode2, 7)
	linkNode3 := MergeTwoLists(linkNode, linkNode2)
    for linkNode3 != nil {
		fmt.Printf("%d ", linkNode3.Val)
		linkNode3 = linkNode3.Next
	}

    fmt.Println("Link test")
}


func TestLinkInsertHead(t *testing.T) {
	linkNode := &LinkNode{Val: 1}
	linkNode = linkNode.InsertHead(linkNode, 4)
	linkNode = linkNode.InsertHead(linkNode, 3)
	linkNode = linkNode.InsertHead(linkNode, 4)

	linkNode = linkNode.Reverse(linkNode)
	for linkNode != nil {
		fmt.Printf("%d ", linkNode.Val)
		linkNode = linkNode.Next
	}
	linkNode = nil // release memory
    fmt.Println("Link test")
}