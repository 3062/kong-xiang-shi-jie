package board

import (
	"kong-xiang-shi-jie/pkg/option"

	"github.com/hajimehoshi/ebiten/v2"
)

var currentImage *ebiten.Image
var defaultOpt *ebiten.DrawImageOptions // 一般不设置此选项

// 渲染帧调用，获取当前图像
func GetDrawImage() (*ebiten.Image, *ebiten.DrawImageOptions) {
	return currentImage, defaultOpt
}

// 逻辑帧调用，重置 image
func InitBoard() {
	currentImage = ebiten.NewImage(option.Width, option.Height)
}

// 逻辑帧调用，逐层叠加
func SetImage(image *ebiten.Image, opt *ebiten.DrawImageOptions) {
	currentImage.DrawImage(image, opt)
}
