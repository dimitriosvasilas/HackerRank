package algo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinaryTreeSer(t *testing.T) {
	root, _ := newNodeStrValue("root")
	n, _ := newNodeStrValue("left")
	root.left = n
	n, _ = newNodeStrValue("left.left")
	root.left.left = n
	n, _ = newNodeStrValue("right")
	root.right = n

	assert.Equal(t, "left.left", deserialize(serialize(root)).left.left.value.(nodeStrValue).v, "")
}
