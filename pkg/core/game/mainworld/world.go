// Copyright 2016 The Ebiten Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mainworld

import (
	"image"
	"kong-xiang-shi-jie/pkg/core/controller"
	"kong-xiang-shi-jie/pkg/player"
	"kong-xiang-shi-jie/pkg/types"
)

func init() {
	err := controller.RegisterGame(GameID, &defaultGame)
	if err != nil {
		panic(err)
	}
}

// Game represents a game state.
type Game struct {
	Player      player.Player
	Background  image.Image
	Map         []types.BlockType
	Animals     []types.Animal
	Environment Environment
}

var defaultGame = Game{}
var GameID types.GameID = 1

func GetGame() Game {
	return defaultGame
}

func (g *Game) initialization() {
}

func (g *Game) Update() error {
	g.Player.Update()
	for _, animal := range g.Animals {
		animal.Update()
	}
	g.Environment.Update()

	return nil
}

func (g *Game) Load() error {
	return nil
}
