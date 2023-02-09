package vector

import (
	"math"

	"kong-xiang-shi-jie/tool"
)

type Vector2[T tool.Number] struct {
	X, Y T
}

func (v Vector2[T]) Add(vector Vector2[T]) Vector2[T] {
	v.X += vector.X
	v.Y += v.Y
	return v
}

func (v Vector2[T]) Sub(vector Vector2[T]) Vector2[T] {
	v.X -= vector.X
	v.Y -= v.Y
	return v
}

func (v Vector2[T]) Len() T {
	return T(math.Sqrt(float64(v.X*v.X + v.Y*v.Y)))
}

func (v Vector2[T]) Len2() T {
	return v.X*v.X + v.Y*v.Y
}
