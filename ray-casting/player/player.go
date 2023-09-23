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

func (p *Player) Update() {
	p.move()
	p.draw()
}

func (p *Player) move() {
	if rl.IsKeyDown(rl.KeyA) {
		p.Position.X -= 5
	}
	if rl.IsKeyDown(rl.KeyD) {
		p.Position.X += 5
	}
	if rl.IsKeyDown(rl.KeyW) {
		p.Position.Y -= 5
	}
	if rl.IsKeyDown(rl.KeyS) {
		p.Position.Y += 5
	}
	if rl.IsKeyDown(rl.KeyLeft) {
		p.Angle -= 0.0625
	}
	if rl.IsKeyDown(rl.KeyRight) {
		p.Angle += 0.0625
	}
}

func (p *Player) draw() {
	angle := p.Angle / 2
	for i := 0; i < len(p.rays); i++ {
		x := float64(p.Position.X) + 500*math.Cos(float64(angle))
		y := float64(p.Position.Y) + 500*math.Sin(float64(angle))

		rl.DrawLine(int32(p.Position.X), int32(p.Position.Y), int32(x), int32(y), color.RGBA{255, 255, 255, 255})
		angle += 0.0625
	}

	rl.DrawCircle(int32(p.Position.X), int32(p.Position.Y), 10, color.RGBA{255, 255, 255, 255})
}
