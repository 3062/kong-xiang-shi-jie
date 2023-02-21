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

type TreeType int

func NewTree() *Tree {
	return nil
}

func NewRectangle(length, width, height int) *Tree {
	return nil
}

func NewDiamond(param1, param2, height int) *Tree {
	return nil
}

func NewEllipse(param1, param2, height int) *Tree {
	return nil
}
