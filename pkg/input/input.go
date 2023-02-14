package input

import "github.com/hajimehoshi/ebiten/v2"

type Input struct {
	*Mouse
	*Keyboard
}

var defaultInput = Input{
	Mouse:    &Mouse{},
	Keyboard: &Keyboard{make(map[ebiten.Key]Continued)},
}

func Update() {
	defaultInput.Mouse.Update()
	defaultInput.Keyboard.Update()
}

func GetMouse() *Mouse {
	return defaultInput.Mouse
}

func GetKeyboard() *Keyboard {
	return defaultInput.Keyboard
}

// 各个模块将自身需要记录的按键注册至 Keyboard
func RegisterKey(key ebiten.Key) {
	defaultInput.Keyboard.Keys[key] = Continued{}
}

// 效果等同于获取 Keyboard 再判断
func KeyIsClick(key ebiten.Key) bool {
	return defaultInput.Keyboard.Keys[key].IsClick()
}

type Continued struct {
	Press   int // 记录按键持续时长（逻辑帧数）
	Release int // 松开时，记录 Press
}

// 更新时传入当前帧状态
func (c *Continued) Update(state bool) {
	if state {
		c.Press += 1
		c.Release = 0
	} else {
		c.Release = c.Press
		c.Press = 0
	}
}

// 检测按下按键的第一帧
func (c Continued) IsClick() bool {
	if c.Press == 1 {
		return true
	}
	return false
}

// 检测持续按下按键
func (c Continued) IsPress() bool {
	if c.Press > 0 {
		return true
	}
	return false
}

// 检测松开按键的第一帧
func (c Continued) IsRelease() bool {
	if c.Release > 0 {
		return true
	}
	return false
}

func (c Continued) GetPress() int {
	return c.Press
}

func (c Continued) GetRelease() int {
	return c.Release
}
