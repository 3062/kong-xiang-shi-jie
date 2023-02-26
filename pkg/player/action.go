package player

import "kong-xiang-shi-jie/pkg/types"

type ActionFunc func(*Player)

func NoneAction(*Player) {}

func NewDestroyAction(intensity float64) ActionFunc {
	return func(p *Player) {

	}
}

func NewAttackAction() ActionFunc {
	return func(p *Player) {

	}
}

func NewPlaceAction(types.EntityType) ActionFunc {
	return func(p *Player) {

	}
}

func NewRunAction(speed float64) ActionFunc {
	return func(p *Player) {

	}
}

func NewRunAction(speed float64) ActionFunc {
	return func(p *Player) {

	}
}
