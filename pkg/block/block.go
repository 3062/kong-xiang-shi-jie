package block

import (
	"kong-xiang-shi-jie/pkg/types"
	"kong-xiang-shi-jie/tool/vector"
)

type Block struct {
	Position   vector.Vector2[int]
	EntityType types.EntityType
	Destroyed  func() types.EntityType
}

func (b *Block) GetPosition() vector.Vector2[int] {
	return b.Position
}

func (b *Block) GetEntityType() types.EntityType {
	return b.EntityType
}
