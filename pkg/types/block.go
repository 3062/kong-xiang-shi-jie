package types

import (
	"kong-xiang-shi-jie/tool/queue"
	"kong-xiang-shi-jie/tool/vector"
)

// 构成世界的方块实体
type Block interface {
	Entity
	Destroyed()
	BuildStuff(vector.Vector2[float64])

	// 统一结算队列，在逻辑帧 update 最后执行，主要用于大范围方块破坏时，等待全部结算后再执行一些判断行为。
	// 后续可能放到公共区域，而不是存储在 Block 中
	SettlementQueue(queue.Queue[func()])
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
