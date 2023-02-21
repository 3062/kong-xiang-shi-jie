package block

import (
	"kong-xiang-shi-jie/pkg/types"
	"kong-xiang-shi-jie/tool/queue"
	"kong-xiang-shi-jie/tool/vector"
)

type Block struct {
	Position  vector.Vector2[int]
	BlockType types.BlockType
}

func (b *Block) GetPosition() vector.Vector2[int] {
	return b.Position
}

func (b *Block) GetBlockType() types.BlockType {
	return b.BlockType
}

func (b *Block) Destroyed(q *queue.Queue[func()]) types.BlockType {
	return b.BlockType
}
