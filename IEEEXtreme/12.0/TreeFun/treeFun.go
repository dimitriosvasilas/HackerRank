package main

import (
	"errors"
	"fmt"
)

var preIndex int

type treeNode struct {
	label int
	left  *treeNode
	right *treeNode
	value int
}

func newNode(label int) *treeNode {
	return &treeNode{
		label: label,
		left:  nil,
		right: nil,
		value: 0,
	}
}

func readN(all []int, i, n int) {
	if n == 0 {
		return
	}
	if m, err := scan(&all[i]); m != 1 {
		panic(err)
	}
	readN(all, i+1, n-1)
}

func scan(a *int) (int, error) {
	return fmt.Scan(a)
}

func print(node *treeNode) {
	if node == nil {
		return
	}
	fmt.Println(node.label)
	print(node.left)
	print(node.right)
}

func main() {
	nm := make([]int, 2)
	readN(nm, 0, 2)
	N := nm[0]
	M := nm[1]
	fmt.Println(N, M)
	tree := make(map[int]*treeNode)
	for i := 0; i < N; i++ {
		uv := make([]int, 2)
		readN(uv, 0, 2)
		if _, ok := tree[uv[0]]; !ok {
			tree[uv[0]] = newNode(uv[0])
		}
		if _, ok := tree[uv[1]]; !ok {
			tree[uv[1]] = newNode(uv[1])
		}
		if tree[uv[0]].left == nil {
			tree[uv[0]].left = tree[uv[1]]
		} else if tree[uv[0]].right == nil {
			tree[uv[0]].right = tree[uv[1]]
		} else {
			panic(errors.New("tree not binary"))
		}
		fmt.Println(tree)
	}
	print(tree[0])
}
