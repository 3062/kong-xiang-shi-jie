package option

import (
	"image"
	"kong-xiang-shi-jie/pkg/button"
)

type PanelID int

const (
	StartID  PanelID = iota // 开始界面
	LoadID                  // 加载界面
	OptionID                // 选项界面
	VolumeID                // 音量界面
)

var currentPanel *OptionPanel
var PanelSet map[PanelID]*OptionPanel

func RegisterPanel(id PanelID, panel *OptionPanel) {
	PanelSet[id] = panel
}

func SwitchPanel(id PanelID) {
	currentPanel = PanelSet[id]
}

func Update() (err error) {
	return currentPanel.Update()
}

type OptionPanel struct {
	background image.Image
	buttons    []*button.Button
}

func (c *OptionPanel) Update() (err error) {
	for _, button := range c.buttons {
		err = button.Update()
	}
	return
}

func (c *OptionPanel) SetBackground(img image.Image) {
	c.background = img
}

func (c *OptionPanel) SetButton(b *button.Button) {
	c.buttons = append(c.buttons, b)
}
