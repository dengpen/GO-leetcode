package main

import (
	"fmt"
	"math"
)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	return height(root) != -1
}

func height(root *TreeNode) float64 {
	if root == nil {
		return 0
	}

	left := height(root.Left)

	if left == -1 {
		return -1
	}

	right := height(root.Right)
	if right == -1 {
		return -1
	}

	fmt.Println(root.Val)
	if math.Abs(left-right) < 2 {
		return math.Max(left, right) + 1
	} else {
		return -1
	}

}

func isBalanced1(root *TreeNode) bool {
	return false
}

func createTree(v int) *TreeNode {
	return &TreeNode{v, nil, nil}
}

func generateTree(node *TreeNode, arr []int) *TreeNode {

	if len(arr) == 1 {
		node.Left = &TreeNode{arr[0], nil, nil}
		return node
	}

	if len(arr) == 2 {
		node.Left = &TreeNode{arr[0], nil, nil}
		node.Right = &TreeNode{arr[1], nil, nil}
		return node
	}

	if arr[0] != 0 {
		node.Left = generateTree(&TreeNode{arr[0], nil, nil}, arr[2:])
	} else {
		node.Left = nil
	}

	if arr[1] != 0 {
		node.Right = generateTree(&TreeNode{arr[1], nil, nil}, arr[2:])
	} else {
		node.Right = nil
	}

	return node
}

func main() {
	tmpArr := []int{1, 2, 2, 3, 3, 0, 0, 4, 4}

	node := createTree(tmpArr[0])

	tmpArr = tmpArr[1:]

	node = generateTree(node, tmpArr)

	// leet函数
	bool := isBalanced(node)

	fmt.Println(bool)
}
