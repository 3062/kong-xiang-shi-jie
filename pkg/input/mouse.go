package input

import (
	"kong-xiang-shi-jie/tool/vector"

	"github.com/hajimehoshi/ebiten/v2"
)

type mouseState int

const (
	release mouseState = iota
	relaxation
	click
	press
)

type Mouse struct {
	LeftButton   mouseState
	RightButton  mouseState
	MiddleButton mouseState
	Roller       float64
	Position     vector.Vector2[int]
}

func (m *Mouse) Update() {
	m.Position.X, m.Position.Y = ebiten.CursorPosition()

	m.LeftButton = ButtonState(m.LeftButton, ebiten.MouseButtonLeft)
	m.RightButton = ButtonState(m.RightButton, ebiten.MouseButtonRight)
	m.MiddleButton = ButtonState(m.MiddleButton, ebiten.MouseButtonMiddle)
}

func ButtonState(current mouseState, button ebiten.MouseButton) mouseState {
	if ebiten.IsMouseButtonPressed(button) {
		if current == press || current == click {
			return press
		} else {
			return click
		}
	} else {
		if current == press || current == click {
			return release
		} else {
			return relaxation
		}
	}
}
