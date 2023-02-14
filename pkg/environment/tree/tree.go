package tree

import (
	"kong-xiang-shi-jie/pkg/block"
	"kong-xiang-shi-jie/tool/vector"
)

type Tree struct {
	root   vector.Vector2[int]
	trunks []block.Block
	leafs  []block.Block
}

func NewTree() *Tree {
	return nil
}
