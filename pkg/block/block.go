package block

import (
	"kong-xiang-shi-jie/pkg/types"
	"kong-xiang-shi-jie/tool/queue"
	"kong-xiang-shi-jie/tool/vector"
)

type Block struct {
	Position   vector.Vector2[int]
	EntityType types.EntityType
}

func (b *Block) GetPosition() vector.Vector2[int] {
	return b.Position
}

func (b *Block) GetEntityType() types.EntityType {
	return b.EntityType
}

func (b *Block) Destroyed(q *queue.Queue[func()]) types.EntityType {
	return b.EntityType
}
