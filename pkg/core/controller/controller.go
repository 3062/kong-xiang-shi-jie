package controller

import (
	"errors"
	"math/rand"
	"time"

	"kong-xiang-shi-jie/pkg/board"
	"kong-xiang-shi-jie/pkg/input"
	"kong-xiang-shi-jie/pkg/option"
	"kong-xiang-shi-jie/pkg/types"

	"github.com/hajimehoshi/ebiten/v2"
)

func init() {
	rand.Seed(time.Now().UnixNano())
	defaultController = MainController{}
	input.RegisterKey(ebiten.KeyEscape)
}

type MainController struct {
	Current        types.Game
	SubControllers []types.Game
}

var defaultController MainController

func GetController() *MainController {
	return &defaultController
}

// 由游戏“世界”层级注册至主控制器，世界编号禁止重复
func RegisterGame(ID types.GameID, game types.Game) error {
	if int(ID) > len(defaultController.SubControllers) {
		return errors.New("error state")
	}
	if defaultController.SubControllers[ID] != nil {
		return errors.New("repeat state")
	}
	defaultController.SubControllers[ID] = game
	return nil
}

// 显示像素大小
func (c *MainController) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return option.Width, option.Height
}

// 逻辑帧更新，处于选项界面则由 option 负责更新，处于游戏界面则由当下游戏世界负责更新
func (c *MainController) Update() error {
	input.Update()

	// 进入选项页面
	if input.KeyIsClick(ebiten.KeyEscape) && !option.IsOption {
		option.IsOption = true
	}

	if option.IsOption {
		return option.Update()
	} else {
		return c.Current.Update()
	}
}

// 渲染帧更新
func (c *MainController) Draw(screen *ebiten.Image) {
	screen.DrawImage(board.GetDrawImage())
}
