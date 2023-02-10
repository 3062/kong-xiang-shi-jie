package input

import "github.com/hajimehoshi/ebiten/v2"

type Keyboard struct {
	Keys []Key
}

type Key struct {
	click bool
	press int
}

func (m *Keyboard) Update() {
	keys := ebiten.AppendInputChars(nil)
	for _, key := range keys {

	}
}
