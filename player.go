package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Player struct {
	sprite                rl.Texture2D
	src                   rl.Rectangle
	dest                  rl.Rectangle
	isMoving              bool
	dir                   int
	up, down, right, left bool
	frame                 int
	speed                 float32
}

func NewPlayer() *Player {
	return &Player{
		sprite:   rl.LoadTexture("res/Characters/BasicCharakterSpritesheet.png"),
		speed:    1.0,
		src:      rl.NewRectangle(0, 0, 48, 48),
		dest:     rl.NewRectangle(200, 200, 100, 100),
		isMoving: false,
		dir:      0,
		up:       false,
		down:     false,
		right:    false,
		left:     false,
		frame:    0,
	}
}

func (p *Player) Input() {
	if (rl.IsKeyDown(rl.KeyLeftShift) || rl.IsKeyDown(rl.KeyRightShift)) && p.speed < 5.5 {
		p.speed += 0.25
	} else if p.speed > 1.0 {
		p.speed -= 0.125
	}
	if rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyUp) {
		p.isMoving, p.up, p.dir = true, true, 1
	}
	if rl.IsKeyDown(rl.KeyS) || rl.IsKeyDown(rl.KeyDown) {
		p.isMoving, p.down, p.dir = true, true, 0
	}
	if rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight) {
		p.isMoving, p.right, p.dir = true, true, 3
	}
	if rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft) {
		p.isMoving, p.left, p.dir = true, true, 2
	}
}

func (p *Player) Update() {
	p.src.X = p.src.Width * float32(p.frame)

	if p.isMoving {
		if p.up {
			p.dest.Y -= p.speed
		}
		if p.down {
			p.dest.Y += p.speed
		}
		if p.right {
			p.dest.X += p.speed
		}
		if p.left {
			p.dest.X -= p.speed
		}
	}
	// change player frame if moving or idle in 1/2 seconds
	if (p.isMoving && frameCount%8 == 1) || (!p.isMoving && frameCount%30 == 1) {
		p.frame++
	}

	if p.frame > 3 || (!p.isMoving && p.frame > 1) {
		p.frame = 0
	}
	p.src.X = p.src.Width * float32(p.frame)
	p.src.Y = p.src.Height * float32(p.dir)
	p.isMoving, p.up, p.down, p.right, p.left = false, false, false, false, false
}

func (p *Player) Draw() {
	rl.DrawTexturePro(p.sprite, p.src, p.dest, rl.NewVector2(p.dest.Width, p.dest.Height), 0.0, rl.White)
}

func (p *Player) Close() error {
	rl.UnloadTexture(p.sprite)
	return nil
}
