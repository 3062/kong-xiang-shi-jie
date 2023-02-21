package types

import "kong-xiang-shi-jie/tool/queue"

type BlockType int32

const NoneBlock BlockType = 0      //无方块
const IntervalSize BlockType = 100 // 类型方块区间大小，1～100

const (
	TreeBegin BlockType = IntervalSize*iota + 1
	SoilBegin
	LeafBegin
)

const (
	TreeEnd BlockType = IntervalSize*iota + IntervalSize
	SoilEnd
	LeafEnd
)

type Block interface {
	GetBlockType() BlockType
	Destroyed()
	SettlementQueue(queue.Queue[func()]) // 统一结算队列，在逻辑帧 update 最后执行，主要用于大范围方块破坏时，等待全部结算后再执行一些判断行为
}
