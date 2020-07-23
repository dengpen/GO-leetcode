package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归自下而上查找
func generateTrees(n int) []*TreeNode {

	if n == 0 {
		return nil
	}

	return getTree(1, n)
}

func getTree(start, end int) []*TreeNode {
	retArr := make([]*TreeNode, 0)

	// break func
	if start > end {
		return []*TreeNode{nil}
	}

	// left node
	for i := start; i <= end; i++ {
		leftArr := getTree(start, i-1)
		rightArr := getTree(i+1, end)
		// create current node
		for _, left := range leftArr {
			for _, right := range rightArr {
				curNode := &TreeNode{i, nil, nil}
				curNode.Left = left
				curNode.Right = right
				retArr = append(retArr, curNode)
			}
		}
	}

	return retArr
}

// 迭代自上而下添加
func generateTrees1(n int) []*TreeNode {

	// return treeNode
	retNode = make([]*TreeNode, 0)

	if n == 0 {
		return nil
	}

	// save status Node
	

	for i := 1; i <= n; i++ {

	}

	return getTree(1, n)
}

func main() {
	ret := generateTrees(3)
	fmt.Printf("arr: %v", ret[1])
}
