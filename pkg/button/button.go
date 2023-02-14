package button

import (
	"kong-xiang-shi-jie/pkg/input"
	"kong-xiang-shi-jie/tool/figure"
)

type Button struct {
	Rect      figure.Rectangle[int]
	Text      string
	MouseDown bool
	onPressed func() error
}

// 按键触发函数
func (b *Button) SetOnPressed(f func() error) {
	b.onPressed = f
}

func (b *Button) Update() error {
	mouse := input.GetMouse()
	if !mouse.LeftButton.IsRelease() {
		return nil
	}

	p := mouse.Position

	if b.Rect.IsPointIn(p) {
		return b.onPressed()
	}
	return nil
}
