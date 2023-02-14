package player

import (
	"kong-xiang-shi-jie/pkg/config"
	"kong-xiang-shi-jie/pkg/input"
	"kong-xiang-shi-jie/pkg/types"

	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	HP              float64
	Image           []byte
	Buff            []types.Buff
	LeftButtonFunc  ActionFunc
	RightButtonFunc ActionFunc
	RunFunc         ActionFunc
	PlayerStates    []types.PlayerState
}

func NewPlay(cfg *config.Player) Player {
	return Player{}
}

func (p *Player) Update() {
	mouse := input.GetMouse()
	keyboard := input.GetKeyboard()
	if mouse.LeftButton.IsClick() {
		p.LeftButtonFunc(p)
	}

	if mouse.RightButton.IsClick() {
		p.RightButtonFunc(p)
	}

	if keyboard.IsClick(ebiten.Key0) {

	}
}
