package types

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
