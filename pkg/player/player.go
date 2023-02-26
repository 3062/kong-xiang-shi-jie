package player

import (
	"kong-xiang-shi-jie/pkg/config"
	"kong-xiang-shi-jie/pkg/input"
	"kong-xiang-shi-jie/pkg/types"
	"kong-xiang-shi-jie/tool/vector"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/solarlune/resolv"
)

type Player struct {
	HP              float64
	Image           []byte
	Buff            []types.Buff
	LeftButtonFunc  ActionFunc
	RightButtonFunc ActionFunc
	RunFunc         ActionFunc
	PlayerStates    []types.PlayerState
	Collider        *resolv.Object
	Position        vector.Vector2[int]
	Speed           int
}

func NewPlay(cfg *config.Player) *Player {
	p := &Player{
		HP:           cfg.HP,
		Buff:         cfg.Buff,
		PlayerStates: cfg.PlayerStates,
		Position:     cfg.Position,
	}

	p.Collider = resolv.NewObject()

	return p
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

	if keyboard.IsClick(ebiten.KeyRight) {

	}
}

func DefaultLeft() ActionFunc {
	return NewDestroyAction(10)
}

func DefaultRight() ActionFunc {
	return nil
}

func DefaultRun(p *Player) {

}
