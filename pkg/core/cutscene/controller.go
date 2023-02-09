package cutscene

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

type CutsceneController struct {
	Current           ebiten.Game
	Games             []ebiten.Game
	ParentSwitchState func(types.ControllerState)
}

var defaultController = CutsceneController{}

func GetController() *CutsceneController {
	return &defaultController
}

func RegisterSubController(state types.ControllerState, controller types.Controller) (func(types.ControllerState), error) {
	if int(state) > len(defaultController.Games) {
		return nil, errors.New("register cutscene error state")
	}
	if defaultController.Games[state] != nil {
		return nil, errors.New("register cutscene repeat state")
	}
	defaultController.Games[state] = controller
	return defaultController.SwitchState, nil
}

func (c *CutsceneController) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func (c *CutsceneController) Update() error {
	return c.Current.Update()
}

func (c *CutsceneController) Draw(screen *ebiten.Image) {
	c.Current.Draw(screen)
}

func (c *CutsceneController) SwitchState(state types.ControllerState) {
	if types.StateJudge(state, types.Option) {
		c.Current = c.Games[state]
	} else {
		c.ParentSwitchState(state)
	}
}
