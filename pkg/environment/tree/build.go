package tree

import (
	"kong-xiang-shi-jie/pkg/types"
	"math/rand"
)

func NewTree() *Tree {
	treeType := rand.Int() % types.TreeTotal

	switch treeType {

	}

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
