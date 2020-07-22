package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 如果二叉树每个节点都具有相同的值，那么该二叉树就是单值二叉树。

// 只有给定的树是单值二叉树时，才返回 true；否则返回 false。

// 输入：[1,1,1,1,1,null,1]
// 输出：true
func isUnivalTree(root *TreeNode) bool {

	//
	if root == nil {
		return true
	}

	if root.Left != nil && root.Val != root.Left.Val {
		return false
	}

	if root.Right != nil && root.Val != root.Right.Val {
		return false
	}

	if root.Left == nil && root.Right == nil {
		return true
	}

	return isUnivalTree(root.Left) && isUnivalTree(root.Right)
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
	}else{
		node.Left = nil
	}

	if arr[1] != 0 {
		node.Right = generateTree(&TreeNode{arr[1], nil, nil}, arr[2:])
	}else{
		node.Right = nil
	}

	return node
}

func main() {
	tmpArr := []int{8,0,8,8,0,9,8}

	node := createTree(tmpArr[0])

	tmpArr = tmpArr[1:]

	node = generateTree(node, tmpArr)

	// leet函数
	bool := isUnivalTree(node)

	fmt.Println(bool)
}
