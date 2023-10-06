package main

import (
	"image/color"
	"math"
	"math/rand"

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

	p := 0.01

	nx := int32(400)
	ny := int32(400)
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

		nx += int32(float32(math.Sin(p+math.Cos(rand.Float64())*2.62)*3.06-math.Sin(p*0.96)*3.65)/1.301 + rand.Float32() - 0.5)
		ny += int32(float32(math.Cos(p+math.Sin(rand.Float64())*1.34)*5.32-math.Cos(p*1.27)*4.21) / 1.23)
		p += 0.01

		rl.EndDrawing()
	}
	rl.CloseWindow()
}
