package figure

import (
	"kong-xiang-shi-jie/tool"
	"kong-xiang-shi-jie/tool/vector"
)

type Rectangle[T tool.Number] struct {
	Min, Max vector.Vector2[T]
}

// 判断点 p 是否在矩阵中，矩阵用左下点 minP，以及右上点 maxP 表示
func (r *Rectangle[T]) IsPointIn(p vector.Vector2[T]) bool {
	if p.X >= r.Min.X && p.Y >= r.Min.Y && p.X <= r.Max.X && p.Y <= r.Max.Y {
		return true
	}
	return false
}
