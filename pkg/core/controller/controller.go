package controller

import (
	"errors"
	"math/rand"
	"time"

	"kong-xiang-shi-jie/pkg/board"
	"kong-xiang-shi-jie/pkg/input"
	"kong-xiang-shi-jie/pkg/types"

	"github.com/hajimehoshi/ebiten/v2"
)

func init() {
	rand.Seed(time.Now().UnixNano())
	defaultController = MainController{}
}

type MainController struct {
	Current        types.Game
	SubControllers []types.Game
}

var defaultController MainController

// NewGame generates a new Game object.
func GetController() *MainController {
	return &defaultController
}

func RegisterGame(state types.GameIden, game types.Game) error {
	if int(state) > len(defaultController.SubControllers) {
		return errors.New("error state")
	}
	if defaultController.SubControllers[state] != nil {
		return errors.New("repeat state")
	}
	defaultController.SubControllers[state] = game
	return nil
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
	screen.DrawImage(board.GetDrawImage())
}
