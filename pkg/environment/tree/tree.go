package tree

import (
	"kong-xiang-shi-jie/pkg/block"
	"kong-xiang-shi-jie/pkg/types"
)

type Tree struct {
	trunk []block.Block
	leafs []block.Block
	shape types.TreeShape
}
