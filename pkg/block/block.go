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

func DeleteByPosition(bl []Block, p vector.Vector2[int]) ([]Block, bool) {
	for i, b := range bl {
		if b.Position.Equal(p) {
			bl = append(bl[0:i], bl[i:]...)
			return bl, true
		}
	}
	return bl, false
}
