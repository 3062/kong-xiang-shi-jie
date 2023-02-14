package option

import (
	"kong-xiang-shi-jie/pkg/button"
	"kong-xiang-shi-jie/pkg/option"
)

func init() {
	option.RegisterPanel(option.StartID, newStartPanel())
	option.SwitchPanel(option.StartID)
}

func newStartPanel() *option.OptionPanel {
	opt := &option.OptionPanel{}
	opt.SetBackground(nil)
	opt.SetButton(newNGameButton())
	opt.SetButton(newContinueButton())
	opt.SetButton(newLoadButton())
	opt.SetButton(newOptionButton())
	opt.SetButton(newExitButton())

	return opt
}

func newNGameButton() *button.Button {
	b := &button.Button{}
	return b
}

func newContinueButton() *button.Button {
	b := &button.Button{}
	return b
}

func newLoadButton() *button.Button {
	b := &button.Button{}
	return b
}

func newOptionButton() *button.Button {
	b := &button.Button{}
	return b
}

func newExitButton() *button.Button {
	b := &button.Button{}
	return b
}
