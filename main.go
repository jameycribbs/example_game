package main

import (
	"github.com/jameycribbs/game_manager"
	"github.com/veandco/go-sdl2/sdl"
)

const FPS = 60

const DELAY_TIME = 1000 / FPS

func main() {
	var frameStart, frameTime uint32

	game := game_manager.NewGame()

	if game.Setup("Chapter 1", 100, 100, 640, 480, false) {
		for game.Running() {
			frameStart = sdl.GetTicks()

			game.HandleEvents()
			game.Update()
			game.Render()

			frameTime = sdl.GetTicks() - frameStart

			if frameTime < DELAY_TIME {
				sdl.Delay(uint32(DELAY_TIME - frameTime))
			}
		}
	}

	game.Clean()
}
