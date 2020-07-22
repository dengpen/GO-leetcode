package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 迭代自上而下查找
func generateTrees(n int) {

	for i := 1; i <= n; i++ {
		// 创建根节点
		tmpNode := &TreeNode{i, nil, nil}

		retArr := getTree(tmpNode, i, n)
	}

	fmt.Printf("reArr: %v", retArr)
}

func getTree(node *TreeNode, i int, n int) *TreeNode {
	retArr := make([]*TreeNode, 0)

	if i == 0 {
		return nil
	}

	if n == 1 || i >= n {
		return node
	}

	for 

	node.Left = getTree(&TreeNode{i - 1, nil, nil}, i-1, n)

	node.Right = getTree(&TreeNode{i + 1, nil, nil}, i+1, n)

	return node
}

func main() {
	generateTrees(3)
}
