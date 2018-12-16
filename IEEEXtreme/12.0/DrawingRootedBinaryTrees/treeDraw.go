package main

import (
	"errors"
	"fmt"
	"log"
)

var preIndex int

type treeNode struct {
	label     byte
	left      *treeNode
	right     *treeNode
	nodesleft int
	pos       int
}

func newNode(label byte) *treeNode {
	return &treeNode{
		label:     label,
		left:      nil,
		right:     nil,
		nodesleft: 0,
		pos:       0,
	}
}

func nodesLeft(node *treeNode) int {
	if node == nil {
		return 0
	}
	chLeft := nodesLeft(node.left)
	chRight := nodesLeft(node.right)
	node.nodesleft = chLeft
	return 1 + chLeft + chRight
}

func pos(node *treeNode, offset int) {
	if node == nil {
		return
	}
	node.pos = offset + node.nodesleft
	pos(node.left, offset)
	pos(node.right, node.pos+1)

}

func readLine() (string, string, bool) {
	var inF string
	var preF string
	fmt.Scanln(&inF)
	if inF == "" {
		return "", "", false
	}
	fmt.Scanln(&preF)
	return inF, preF, true
}

func search(arr string, start int, end int, value byte) int {
	for i := start; i <= end; i++ {
		if arr[i] == value {
			return i
		}
	}
	return -1
}

func buildTree(inFix string, preFix string, inStart int, inEnd int) *treeNode {
	if inStart > inEnd {
		return nil
	}
	node := newNode(preFix[preIndex])
	preIndex++

	if inStart == inEnd {
		return node
	}

	inIndex := search(inFix, inStart, inEnd, node.label)
	if inIndex == -1 {
		log.Fatal(errors.New("search did not find label in inFix"))
	}

	node.left = buildTree(inFix, preFix, inStart, inIndex-1)
	node.right = buildTree(inFix, preFix, inIndex+1, inEnd)
	return node
}

func print(node *treeNode) {
	if node == nil {
		return
	}
	queue := make([]*treeNode, 0)
	queue = append(queue, node)
	for {
		nodeCount := len(queue)
		if nodeCount == 0 {
			break
		}
		lineOffset := 0
		for nodeCount > 0 {
			n := queue[0]
			for lineOffset < n.pos {
				fmt.Printf(" ")
				lineOffset++
			}
			fmt.Printf("%c", n.label)
			lineOffset++
			queue = queue[1:]
			if n.left != nil {
				queue = append(queue, n.left)
			}
			if n.right != nil {
				queue = append(queue, n.right)
			}
			nodeCount--
		}
		fmt.Printf("\n")
		lineOffset = 0
	}
}

func main() {
	for {
		infix, prefix, f := readLine()
		if f {
			preIndex = 0
			tree := buildTree(infix, prefix, 0, len(infix)-1)
			nodesLeft(tree)
			pos(tree, 0)
			print(tree)
		} else {
			break
		}
	}
}
