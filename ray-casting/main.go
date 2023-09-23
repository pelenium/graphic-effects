package main

import (
	"graphic-effects/ray-casting/player"
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	WIDTH  = 800
	HEIGHT = 800
)

func main() {
	rl.InitWindow(WIDTH, HEIGHT, "ray casting")

	rl.SetTargetFPS(60)

	p := player.Init(WIDTH / 2, HEIGHT / 2, 0)

	
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(color.RGBA{0, 0, 0, 255})

		p.Update()

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
