package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Camera struct {
	cam  rl.Camera2D
	zoom float32
}

func newCamera(pl Player) *Camera {
	return &Camera{
		cam: rl.NewCamera2D(rl.NewVector2(float32(screenWidth/2), float32(screenHeight/2)),
			rl.NewVector2(pl.dest.X-(pl.dest.Width/2), pl.dest.Y-(pl.dest.Height/2)), 0.0, 1.0),
		zoom: 1.0,
	}
}

func (c *Camera) Input() {
	if !rl.IsKeyDown(rl.KeyLeftShift) && rl.IsKeyDown(rl.KeyZ) {
		if c.zoom += 0.1; c.zoom > 3.0 {
			c.zoom = 3.0
		}
	}
	if rl.IsKeyDown(rl.KeyLeftShift) && rl.IsKeyDown(rl.KeyZ) {
		if c.zoom -= 0.1; c.zoom < 0.5 {
			c.zoom = 0.5
		}
	}
}

func (c *Camera) Update() {
	c.cam.Target = rl.NewVector2(player.dest.X-(player.dest.Width/2), player.dest.Y-(player.dest.Height/2))
	c.cam.Zoom = c.zoom
}
