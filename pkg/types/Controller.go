package types

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type ControllerState int

const (
	End ControllerState = iota
	Cutscene
	Game
	Option
)

type Controller interface {
	ebiten.Game
	SwitchState(ControllerState)
}

func StateJudge(target, current ControllerState) bool {
	return true
}
