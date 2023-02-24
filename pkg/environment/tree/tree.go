package tree

import (
	"kong-xiang-shi-jie/pkg/block"
	"kong-xiang-shi-jie/pkg/types"
	"kong-xiang-shi-jie/tool/vector"
)

type Tree struct {
	root           vector.Vector2[int]
	trunk          []block.Block
	leafs          []block.Block
	shape          types.TreeShape
	age            int
	existDestroyed bool
}

func (t *Tree) Update() {

	switch t.age {
	case 100:

	}
}

func (t *Tree) Destroyed(position vector.Vector2[int]) {

}

func (t *Tree) LeafDrop() {

}

func (t *Tree) Grow() {

}
