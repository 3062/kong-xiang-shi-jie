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

	Position vector.Vector2[int]
	Shift    vector.Vector2[int]
	IsMove   bool
}

func (m *Mouse) Update() {
	m.LeftButton = ButtonState(m.LeftButton, ebiten.MouseButtonLeft)
	m.RightButton = ButtonState(m.RightButton, ebiten.MouseButtonRight)
	m.MiddleButton = ButtonState(m.MiddleButton, ebiten.MouseButtonMiddle)

	tempVec := vector.NewVector2[int](ebiten.CursorPosition())
	if m.Position.Equal(tempVec) {
		m.IsMove = false
	} else {
		m.IsMove = true
	}
	m.Shift = m.Position.Sub(tempVec)
	m.Position = tempVec
}

func (m *Mouse) IsPressLeft() bool {
	return m.LeftButton == click || m.LeftButton == press
}

func (m *Mouse) IsPressRight() bool {
	return m.RightButton == click || m.RightButton == press
}

func (m *Mouse) IsPressMiddle() bool {
	return m.MiddleButton == click || m.MiddleButton == press
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
