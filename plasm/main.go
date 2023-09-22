package main

import (
	"image/color"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	width  = 400
	height = 400
)

func main() {
	rl.InitWindow(width, height, "plasm")

	v := float64(0)

	val := float64(0)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(color.RGBA{0, 0, 0, 255})

		for x := float64(0); x < width; x++ {
			for y := float64(0); y < height; y++ {
				val = 0.5*math.Sin(v+1.24*math.Sin(x*0.3+y*0.1)+math.Sin(x*0.02+y*0.37)+3*math.Sin(x*0.15+y*0.08)+1.8*math.Sin(x*0.139+y*0.265)) + 0.5
				if val >= 0 && val < 0.25 {
					rl.DrawPixel(int32(x), int32(y), color.RGBA{64, 64, 64, 255})
				} else if val >= 0.25 && val < 0.5 {
					rl.DrawPixel(int32(x), int32(y), color.RGBA{128, 128, 128, 255})
				} else if val >= 0.5 && val < 0.75 {
					rl.DrawPixel(int32(x), int32(y), color.RGBA{192, 192, 192, 255})
				} else {
					rl.DrawPixel(int32(x), int32(y), color.RGBA{255, 255, 255, 255})
				}
			}
		}

		rl.EndDrawing()

		v += 0.3
	}

	rl.CloseWindow()
}
