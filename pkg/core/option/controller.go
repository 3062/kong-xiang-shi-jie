package option

import (
	"errors"

	"kong-xiang-shi-jie/pkg/core/start"
	"kong-xiang-shi-jie/pkg/types"

	"github.com/hajimehoshi/ebiten/v2"
)

func init() {
	callback, err := start.RegisterSubController(types.Option, &defaultController)
	if err != nil {
		panic(err)
	}
	defaultController.ParentSwitchState = callback
}

type OptionController struct {
	Current           ebiten.Game
	Games             []ebiten.Game
	ParentSwitchState func(types.ControllerState)
}

var defaultController = OptionController{Games: make([]ebiten.Game, 10)}

func GetController() *OptionController {
	return &defaultController
}

func RegisterSubController(state types.ControllerState, controller types.Controller) (func(types.ControllerState), error) {
	if int(state) > len(defaultController.Games) {
		return nil, errors.New("register option error state")
	}
	if defaultController.Games[state] != nil {
		return nil, errors.New("register option repeat state")
	}
	defaultController.Games[state] = controller
	return defaultController.SwitchState, nil
}

func (c *OptionController) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func (c *OptionController) Update() error {
	return c.Current.Update()
}

func (c *OptionController) Draw(screen *ebiten.Image) {
	c.Current.Draw(screen)
}

func (c *OptionController) SwitchState(state types.ControllerState) {
	if types.StateJudge(state, types.Option) {
		c.Current = c.Games[state]
	} else {
		c.ParentSwitchState(state)
	}
}
