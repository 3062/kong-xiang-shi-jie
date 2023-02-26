package config

import (
	"kong-xiang-shi-jie/pkg/types"
	"kong-xiang-shi-jie/tool/vector"
)

type Conifg struct {
}

type Player struct {
	HP           float64
	Buff         []types.Buff
	PlayerStates []types.PlayerState
	Position     vector.Vector2[int]
}

type Option struct {
}

type World struct {
}
