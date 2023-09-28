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

	rays := [100]ray{}

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
		p.Angle -= 0.015625
	}
	if rl.IsKeyDown(rl.KeyRight) {
		p.Angle += 0.015625
	}
}

func (p *Player) draw(gameMap []string) {
	angle := p.Angle / 2
	for i := 0; i < len(p.rays); i++ {
		for c := float64(0); c < 800; c += 1 {
			x := float64(p.Position.X) + c*math.Cos(float64(angle))
			y := float64(p.Position.Y) + c*math.Sin(float64(angle))

			if string(gameMap[int(y)/20][int(x)/20]) == "#" {
				// ray drawing
				// rl.DrawLine(int32(p.Position.X), int32(p.Position.Y), int32(x), int32(y), color.RGBA{255, 255, 255, 255})
				length := int32(math.Sqrt(math.Pow(float64(p.Position.X)-x, 2) + math.Pow(float64(p.Position.Y)-y, 2)) /* * math.Cos(float64(angle))*/)

				var height int32
				if length <= 50 {
					height = 800
				} else if length >= 200 {
					height = 0
				} else {
					height = 5 * length
				}
				height = max(min(800-length, 800), 0)
				var contrastRatio uint8 = max(min(uint8(255-0.32*float32(length)), 255), 0)

				rl.DrawRectangle(int32(i*8), (800-height)/2, 8, height, color.RGBA{contrastRatio, contrastRatio, contrastRatio, 255})

				break
			}
		}
		angle += 0.015625
	}
}

func ToRadians(angle float64) float64 {
	return angle * (math.Pi / 180.0)
}
