package image

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

// 切分图片，参数
// long: (切分后)单个图片的长
// high: 单个图片的高
// longSpace: 图片间横向间隔
// highSpace: 图片间纵向间隔
// columnNum: 列数
// lineNum: 行数
func SplitImage(img *ebiten.Image, long, high, longSpace, highSpace, columnNum, lineNum int) [][]*ebiten.Image {
	grossLong := long + longSpace
	grosshigh := high + highSpace
	var ret = make([][]*ebiten.Image, lineNum)
	for i := 0; i < lineNum; i++ {
		ret = append(ret, make([]*ebiten.Image, columnNum))
		for j := 0; j < columnNum; j++ {
			temp := img.SubImage(image.Rect(j*(grossLong), i*(grosshigh), j*(grossLong)+long, i*(grosshigh)+high)).(*ebiten.Image)
			ret[i] = append(ret[i], temp)
		}
	}
	return ret
}
