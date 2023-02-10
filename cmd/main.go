package main

import (
	"log"

	_ "kong-xiang-shi-jie/extend"
	"kong-xiang-shi-jie/pkg/core/controller"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game := controller.GetController()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
