package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Camera struct {
	player *Player
	cam    rl.Camera2D
	zoom   float32
}

func NewCamera(pl *Player) *Camera {
	return &Camera{
		player: pl,
		cam: rl.NewCamera2D(rl.NewVector2(float32(screenWidth/2), float32(screenHeight/2)),
			rl.NewVector2(pl.dest.X-(pl.dest.Width/2), pl.dest.Y-(pl.dest.Height/2)), 0.0, 1.0),
		zoom: 1,
	}
}

func (c *Camera) Input() {
	if !rl.IsKeyDown(rl.KeyLeftShift) && rl.IsKeyDown(rl.KeyZ) {
		if c.zoom += 0.1; c.zoom > 4.0 {
			c.zoom = 4.0
		}
	}
	if rl.IsKeyDown(rl.KeyLeftShift) && rl.IsKeyDown(rl.KeyZ) {
		if c.zoom -= 0.1; c.zoom < 0.75 {
			c.zoom = 0.75
		}
	}
}

func (c *Camera) Update() {
	c.cam.Target = rl.NewVector2(c.player.dest.X-(c.player.dest.Width/2), c.player.dest.Y-(c.player.dest.Height/2))
	c.cam.Zoom = c.zoom
}
