package input

import (
	"kong-xiang-shi-jie/tool/vector"

	"github.com/hajimehoshi/ebiten/v2"
)

// Mouse 结构体不保证零值可用时，需要在 input 中初始化
type Mouse struct {
	LeftButton   Continued
	RightButton  Continued
	MiddleButton Continued
	Roller       float64

	Position vector.Vector2[int]
	Shift    vector.Vector2[int]
	IsMove   bool
}

func (m *Mouse) Update() {
	m.LeftButton.Update(ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft))
	m.RightButton.Update(ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight))
	m.MiddleButton.Update(ebiten.IsMouseButtonPressed(ebiten.MouseButtonMiddle))

	tempVec := vector.NewVector2[int](ebiten.CursorPosition())
	if m.Position.Equal(tempVec) {
		m.IsMove = false
	} else {
		m.IsMove = true
	}
	m.Shift = m.Position.Sub(tempVec)
	m.Position = tempVec
}
