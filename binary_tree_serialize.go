package algo

import (
	"errors"
	"math"
	"strings"
)

/*
	Given the root to a binary tree, implement serialize(root), which serializes the tree into a string,
	and deserialize(s), which deserializes the string back into the tree.
*/

var separator = ","

type nodeStrValue struct {
	v string
}

func newNodeStrValue(v string) (*treeNode, error) {
	if strings.Contains(v, separator) {
		return nil, errors.New("nodeValue should not contain ','")
	}
	return newNode(nodeStrValue{v: v}), nil
}

func serialize(n *treeNode) string {
	height := treeHeight(n)
	treeArray := make([]string, int(math.Exp2(float64(height))-1))
	_traverseSerialize(n, 0, treeArray)

	return strings.Join(treeArray, ",")
}

func _traverseSerialize(n *treeNode, currectIndex int, serializeArr []string) {
	if n == nil {
		return
	}

	serializeArr[currectIndex] = n.value.(nodeStrValue).v

	_traverseSerialize(n.left, (2*currectIndex)+1, serializeArr)
	_traverseSerialize(n.right, (2*currectIndex)+2, serializeArr)
}

func deserialize(treeStr string) *treeNode {
	treeArr := strings.Split(treeStr, ",")

	return _buildtree(nil, treeArr, 0)
}

func _buildtree(n *treeNode, treeArr []string, arrIdx int) *treeNode {
	if arrIdx >= len(treeArr) || treeArr[arrIdx] == "" {
		return nil
	}

	n, _ = newNodeStrValue(treeArr[arrIdx])

	n.left = _buildtree(n.left, treeArr, 2*arrIdx+1)
	n.right = _buildtree(n.right, treeArr, 2*arrIdx+2)

	return n
}
