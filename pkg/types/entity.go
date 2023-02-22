package types

import "kong-xiang-shi-jie/tool/vector"

type EntityType int32

const NoneEntity EntityType = 0     //无
const IntervalSize EntityType = 100 // 类型区间大小，1～100

const (
	TreeBegin EntityType = IntervalSize*iota + 1
	SoilBegin
	LeafBegin
)

const (
	TreeEnd EntityType = IntervalSize*iota + IntervalSize
	SoilEnd
	LeafEnd
)

type Entity interface {
	GetEntityType() EntityType
}

// 素材，所有可放在背包中的物品
type Stuff interface {
	Entity
	GetPosition() vector.Vector2[float64]
}
