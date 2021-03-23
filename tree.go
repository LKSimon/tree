package main

import (
	"fmt"
)

type TreeNode struct {
	Value interface{}
	Left  *TreeNode
	Right *TreeNode
}

type Tree struct {
	root *TreeNode
}

func (t *Tree) PreOrderTraversalRec(node *TreeNode) {
	if node == nil {
		return
	}
	fmt.Printf("node value： %q\n", node.Value)
	t.PreOrderTraversalRec(node.Left)
	t.PreOrderTraversalRec(node.Right)
}

func (t *Tree) InOrderTraversalRec(node *TreeNode) {
	if node == nil {
		return
	}
	t.InOrderTraversalRec(node.Left)
	fmt.Printf("node value： %q\n", node.Value)
	t.InOrderTraversalRec(node.Right)
}

func (t *Tree) PostOrderTraversalRec(node *TreeNode) {
	if node == nil {
		return
	}
	t.PostOrderTraversalRec(node.Left)
	t.PostOrderTraversalRec(node.Right)
	fmt.Printf("node value： %q\n", node.Value)
}

func (t *Tree) PreOrderTraversalByStack(root *TreeNode) {
	if root == nil {
		return
	}
	//使用队列模拟栈操作
	stack := make(Stack, 0)
	stack.Push(root)
	for len(stack) > 0 {
		node := stack.Pop()
		fmt.Println(node.Value)//根
		if node.Right != nil {
			stack.Push(node.Right)
		}
		if node.Left != nil {
			stack.Push(node.Left)
		}
	}
}

type Stack []TreeNode

func (s *Stack) Push(node *TreeNode) {
	*s = append(*s, *node)
}

func (s *Stack) Pop() TreeNode {
	length := len(*s)
	v := (*s)[length-1]
	*s = (*s)[:length-1]
	return v
}
