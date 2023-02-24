package types

import (
	"kong-xiang-shi-jie/tool/vector"
)

// 构成世界的方块实体
type Block interface {
	Entity
	Destroyed()
	BuildStuff(vector.Vector2[float64])
}

// 方块素材，可以放入背包
type BlockStuff interface {
	Stuff
	BuildBlock(vector.Vector2[int])
}

// 方块/素材 生成函数，初始化时注册。可避免互相引用的情况出现

var BlockFactory map[EntityType]func(vector.Vector2[int]) Block

func RegisterBlockFactory(t EntityType, build func(vector.Vector2[int]) Block) {
	BlockFactory[t] = build
}

var BlockStuffFactory map[EntityType]func(vector.Vector2[float64]) BlockStuff

func RegisterBlockStuffFactory(t EntityType, build func(vector.Vector2[float64]) BlockStuff) {
	BlockStuffFactory[t] = build
}
