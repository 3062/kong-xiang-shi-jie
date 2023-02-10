package player

import (
	"kong-xiang-shi-jie/pkg/config"
	"kong-xiang-shi-jie/pkg/input"
	"kong-xiang-shi-jie/pkg/types"
)

type Player struct {
	HP    float64
	Image []byte
	Buff  []types.Buff
}

type PlayerConfig struct {
}

func NewPlay(cfg *config.Player) Player {

	return Player{}
}

func (p *Player) Update() {
	mouse := input.GetMouse()
	keyboard := input.GetKeyboard()
	if mouse.IsPressLeft() {

	}
}
