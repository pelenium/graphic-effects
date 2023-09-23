package main

import (
	"image/color"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	width  = 800
	height = 800
)

var (
	circles = [20]circle{}
)

func main() {
	rl.InitWindow(width, height, "spiral")

	rl.SetTargetFPS(60)

	angle := float64(0)
	val := float32(100)

	for i := range circles {
		c := circle{}
		for j := 0; j < 120; j++ {
			c.points = append(c.points, point{color.RGBA{255, 255, 255, 255}, rl.Vector2{X: float32(math.Sin(angle)) * val, Y: float32(math.Cos(angle)) * val}, 2})
			angle += 3
			val += 1.5
		}
		circles[i] = c
	}

	nx := int32(100)
	ny := int32(100)
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(color.RGBA{0, 0, 0, 255})

		for _, c := range circles {
			for _, i := range c.points {
				rl.DrawCircle(int32(i.position.X)+nx, int32(i.position.Y)+ny, i.size, i.col)
			}
		}

		for i := range circles {
			c := circle{}
			for j := 0; j < 120; j++ {
				c.points = append(c.points, point{color.RGBA{255, 255, 255, 255}, rl.Vector2{X: float32(math.Sin(angle)) * val, Y: float32(math.Cos(angle)) * val}, 2})
				angle += 3
				val += 1.5
			}
			circles[i] = c
		}

		if val > 360 {
			val = 0
		}

		rl.EndDrawing()
	}
	rl.CloseWindow()
}
