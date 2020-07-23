package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {

	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
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

	tmpArr1 := []int{1, 2, 3}

	tree1 := createTree(tmpArr1[0])

	tmpArr1 = tmpArr1[1:]

	tree1 = generateTree(tree1, tmpArr1)

	// 树2
	tmpArr2 := []int{1, 2, 2}

	tree2 := createTree(tmpArr2[0])

	tmpArr2 = tmpArr2[1:]

	tree2 = generateTree(tree2, tmpArr2)

	ret := isSameTree(tree1, tree2)

	fmt.Println(ret)
}
