package main

import (
	"fmt"
	"math"
)

type tree struct {
	nodes []treeNode
}

type treeNode struct {
	leftBound  int
	rightBound int
}

func new(from []int) *tree {
	nodes := make([]treeNode, size(len(from)))
	tree := &tree{
		nodes: nodes,
	}
	tree.buildTree(from, 0, 0, len(from)-1)
	return tree
}

func (t *tree) buildTree(from []int, index, lBound, rBound int) {
	fmt.Println("buildTree", index, lBound, rBound)
	if lBound == rBound {
		t.nodes[index].leftBound = lBound
		t.nodes[index].rightBound = rBound
		return
	}
	mid := getMid(lBound, rBound)
	t.buildTree(from, 2*index+1, lBound, mid)
	t.buildTree(from, 2*index+2, mid+1, rBound)

	t.nodes[index].leftBound = t.nodes[2*index+1].leftBound
	t.nodes[index].rightBound = t.nodes[2*index+2].rightBound
}

func getMid(start int, end int) int {
	return start + (end-start)/2
}

func size(arrSize int) int {
	//the size of the tree in 2 * 2^(ceil(log2n)) - 1
	return 1<<uint(math.Ceil(1+math.Log2(float64(arrSize)))) - 1
}

func main() {
	fmt.Println(getMid(0, 10))
	fmt.Println(getMid(5, 10))
	fmt.Println(size(6))

	arr := []int{1, 3, 5, 6}
	t := new(arr)
	fmt.Println(t)
}
