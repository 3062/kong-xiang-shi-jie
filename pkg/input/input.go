package input

type Input struct {
	Mouse
	Keyboard
}

var defaultInput = Input{}

func Update() {
	defaultInput.Mouse.Update()
	defaultInput.Keyboard.Update()
}

func GetMouse() Mouse {
	return defaultInput.Mouse
}

func GetKeyboard() Keyboard {
	return defaultInput.Keyboard
}
