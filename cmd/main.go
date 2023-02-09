package main

import (
	"log"

	"kong-xiang-shi-jie/pkg/core/start"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game := start.GetController()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
