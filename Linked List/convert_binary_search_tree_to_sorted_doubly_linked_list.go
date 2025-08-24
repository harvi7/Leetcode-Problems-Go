package linkedlist

// https://www.youtube.com/watch?v=l1hSUOaXLxc&ab_channel=CrackingFAANG

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

var first *Node
var last *Node

func treeToDoublyList(root *Node) *Node {
	if root == nil {
		return nil
	}

	first, last = nil, nil

	inorderLink(root)

	last.Right = first
	first.Left = last

	return first
}

func inorderLink(node *Node) {
	if node == nil {
		return
	}
	inorderLink(node.Left)
	if last != nil {
		last.Right = node
		node.Left = last
	} else {
		first = node
	}
	last = node
	inorderLink(node.Right)
}
