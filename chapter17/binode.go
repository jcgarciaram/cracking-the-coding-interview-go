package main

import "fmt"

type BiNode struct {
	node1 *BiNode
	node2 *BiNode
	data  int
}

type BiNodeLL struct {
	head *BiNode
	tail *BiNode
}

func (ll *BiNodeLL) print() {
	fmt.Print("forwards: ")
	curr := ll.head
	if curr != nil {
		for {
			fmt.Print(curr.data)
			if curr.node2 == nil {break}
			fmt.Print("->")
			curr = curr.node2
		}
	}
	fmt.Print("\n")
	fmt.Print("backwards: ")
	curr = ll.tail
	if curr != nil {
		for {
			fmt.Print(curr.data)
			if curr.node1 == nil {break}
			fmt.Print("->")
			curr = curr.node1
		}
	}
}
	

type BiNodeTree struct {
	root *BiNode
}

type CurrNodeLL struct {
	node *BiNode
}

func inOrderConvert(root *BiNode, curr *CurrNodeLL, ll *BiNodeLL) {
	if root == nil {
		ll.tail = curr.node
		return
	}
	inOrderConvert(root.node1, curr, ll)
	if curr.node == nil {
		ll.head = root
	} else {
		curr.node.node2 = root
	}
	root.node1 = curr.node
	curr.node = root
	inOrderConvert(root.node2, curr, ll)
}

func convertTreeToLinkendList(tree BiNodeTree) BiNodeLL {
	ll := BiNodeLL{}
	curr := CurrNodeLL{}
	inOrderConvert(tree.root, &curr, &ll)
	return ll
}



func main() {
	b1 := &BiNode{nil, nil, 1}
	b3 := &BiNode{nil, nil, 3}
	b5 := &BiNode{nil, nil, 5}
	b7 := &BiNode{nil, nil, 7}
	
	b4 := &BiNode{b3, b5, 4}
	b2 := &BiNode{b1, b4, 2}
	b8 := &BiNode{b7, nil, 8}
	b6 := &BiNode{b2, b8, 6}
	
	tree := BiNodeTree{b6}
	
	ll := convertTreeToLinkendList(tree)
	
	ll.print()
	
	
}
