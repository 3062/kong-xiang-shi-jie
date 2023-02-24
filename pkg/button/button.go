package button

import (
	"kong-xiang-shi-jie/pkg/input"
	"kong-xiang-shi-jie/tool/figure"
)

type Button struct {
	Rect      figure.Rectangle[int]
	Text      string
	MouseDown bool
	OnPressed func() error
}

type Option struct {
	Rect      figure.Rectangle[int]
	Text      string
	OnPressed func() error
}

// TODO: 判断合法性
func NewButton(opt Option) *Button {
	return &Button{
		Rect:      opt.Rect,
		Text:      opt.Text,
		OnPressed: opt.OnPressed,
	}
}

func (b *Button) Update() error {
	mouse := input.GetMouse()
	if !mouse.LeftButton.IsRelease() {
		return nil
	}

	p := mouse.Position

	if b.Rect.IsPointIn(p) {
		return b.OnPressed()
	}
	return nil
}
