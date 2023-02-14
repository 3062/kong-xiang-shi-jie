package input

import "github.com/hajimehoshi/ebiten/v2"

type Keyboard struct {
	Keys map[ebiten.Key]Continued
}

func (b *Keyboard) Update() {
	// 统计按键数量较多时可以自行实现 IsKeyPressed 优化判断
	for k, c := range b.Keys {
		c.Update(ebiten.IsKeyPressed(k))
	}
}

func (b *Keyboard) IsClick(key ebiten.Key) bool {
	return b.Keys[key].IsClick()
}

func (b *Keyboard) IsPress(key ebiten.Key) bool {
	return b.Keys[key].IsPress()
}

func (b *Keyboard) GetRelease(key ebiten.Key) int {
	return b.Keys[key].GetRelease()
}
