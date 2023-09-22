package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"image/color"
)

type point struct {
	col      color.RGBA
	position rl.Vector2
	size     float32
}
