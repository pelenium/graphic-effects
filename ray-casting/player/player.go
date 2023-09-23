package player

import (
	"image/color"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	Position rl.Vector2
	Angle    float32
	rays     []ray
}

func Init(x, y, angle float32) *Player {
	p := Player{
		Position: rl.Vector2{X: x, Y: y},
		Angle:    angle,
	}

	rays := [25]ray{}

	for i := range rays {
		rays[i] = ray{rl.Vector2{X: x + 500, Y: y}}
	}

	p.rays = rays[:]

	return &p
}

func (p *Player) Update(gameMap []string) {
	p.move(gameMap)
	p.draw(gameMap)
}

func (p *Player) move(gameMap []string) {
	if rl.IsKeyDown(rl.KeyA) {
		if string(gameMap[int(p.Position.Y)/20][int(p.Position.X-2.5)/20]) != "#" {
			p.Position.X -= 2.5
		}
	}
	if rl.IsKeyDown(rl.KeyD) {
		if string(gameMap[int(p.Position.Y)/20][int(p.Position.X+2.5)/20]) != "#" {
			p.Position.X += 2.5
		}
	}
	if rl.IsKeyDown(rl.KeyW) {
		if string(gameMap[int(p.Position.Y-2.5)/20][int(p.Position.X)/20]) != "#" {
			p.Position.Y -= 2.5
		}
	}
	if rl.IsKeyDown(rl.KeyS) {
		if string(gameMap[int(p.Position.Y+2.5)/20][int(p.Position.X)/20]) != "#" {
			p.Position.Y += 2.5
		}
	}
	if rl.IsKeyDown(rl.KeyLeft) {
		p.Angle -= 0.0625
	}
	if rl.IsKeyDown(rl.KeyRight) {
		p.Angle += 0.0625
	}
}

func (p *Player) draw(gameMap []string) {
	angle := p.Angle / 2
	var n uint8
	for i := 0; i < len(p.rays); i++ {
		for c := float64(0); c < 800; c += 1 {
			x := float64(p.Position.X) + c*math.Cos(float64(angle))
			y := float64(p.Position.Y) + c*math.Sin(float64(angle))

			if string(gameMap[int(y)/20][int(x)/20]) == "#" {
				length := int(math.Sqrt(math.Pow(float64(p.Position.X)-x, 2) + math.Pow(float64(p.Position.Y)-y, 2)))

				/* TODO:
				1. Create a formula that counts height
				2. Create a formula that counts contrast ratio
				3. Draw needed rectangle
				*/

				var height int32
				var contrastRatio uint8

				rl.DrawRectangle(int32(i*32), 0, 32, height, color.RGBA{contrastRatio, contrastRatio, contrastRatio, 255})
				break
			}
		}
		angle += 0.0625
	}
}
