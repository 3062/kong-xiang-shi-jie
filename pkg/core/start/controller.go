package start

import (
	"errors"
	"math/rand"
	"time"

	"kong-xiang-shi-jie/pkg/input"
	"kong-xiang-shi-jie/pkg/types"

	"github.com/hajimehoshi/ebiten/v2"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type MainController struct {
	Current        types.Controller
	SubControllers []types.Controller
}

var defaultController = MainController{SubControllers: make([]types.Controller, 5)}

// NewGame generates a new Game object.
func GetController() *MainController {
	return &defaultController
}

func RegisterSubController(state types.ControllerState, controller types.Controller) (func(types.ControllerState), error) {
	if int(state) > len(defaultController.SubControllers) {
		return nil, errors.New("error state")
	}
	if defaultController.SubControllers[state] != nil {
		return nil, errors.New("repeat state")
	}
	defaultController.SubControllers[state] = controller
	return defaultController.SwitchState, nil
}

// Layout implements ebiten.Game's Layout.
func (c *MainController) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

// Update updates the current game state.
func (c *MainController) Update() error {
	input.Update()
	return c.Current.Update()
}

// Draw draws the current game to the given screen.
func (c *MainController) Draw(screen *ebiten.Image) {
	c.Current.Draw(screen)
}

func (c *MainController) SwitchState(state types.ControllerState) {
	c.Current = c.SubControllers[state]
}
