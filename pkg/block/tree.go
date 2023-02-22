package block

import (
	"kong-xiang-shi-jie/pkg/types"
	"kong-xiang-shi-jie/tool/queue"
	"kong-xiang-shi-jie/tool/vector"
)

// 树的非根节点
type Leaf struct {
	Block
	Root *Root
}

func (b *Leaf) Destroyed(q *queue.Queue[func()]) types.EntityType {
	if b.Root.OnceErgodic == false {
		q.PushBack(b.Root.Ergodic)
		b.Root.OnceErgodic = true
	}

	return b.EntityType
}

// 树根节点，树的总节点数较小，使用遍历进行 查找/删除
type Root struct {
	Block
	OnceErgodic bool
	Leafs       []*Leaf
}

func (r *Root) Destroyed(q *queue.Queue[func()]) types.EntityType {
	if r.OnceErgodic == false {
		q.PushBack(r.Ergodic)
		r.OnceErgodic = true
	}
	return r.EntityType
}

func (r *Root) Ergodic() {
	temp := make([]bool, len(r.Leafs))

	for i, t := range temp {
		if !t {
			r.Leafs[i].Destroyed(nil)
		}
	}

	r.OnceErgodic = false
}

func (r *Root) Remove(p vector.Vector2[int]) {
	for i, leaf := range r.Leafs {
		if leaf.Position.Equal(p) {
			r.Leafs = append(r.Leafs[0:i], r.Leafs[i:]...)
			return
		}
	}
}
