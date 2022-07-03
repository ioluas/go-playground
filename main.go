package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	screenWidth  = 1_000
	screenHeight = 480
)

var (
	running bool = true

	backgroundColour rl.Color = rl.NewColor(147, 211, 196, 255)

	frameCount int

	camera  *Camera
	player  *Player
	gameMap *GameMap
	audio   *Audio
)

func input() {
	player.Input()
	camera.Input()
	audio.Input()

	if rl.IsKeyPressed(rl.KeyO) {
		rl.TakeScreenshot("screenshot.png")
	}

	if rl.IsKeyPressed(rl.KeyF) {
		if rl.IsWindowFullscreen() {
			rl.SetWindowSize(screenWidth, screenHeight)
		} else {
			display := rl.GetCurrentMonitor()
			rl.SetWindowSize(rl.GetMonitorWidth(display), rl.GetMonitorHeight(display))
		}
		rl.ToggleFullscreen()
	}
}

func update() {
	running = !rl.WindowShouldClose()
	frameCount++

	player.Update()
	camera.Update()
	audio.Update()
}

func render() {
	rl.BeginDrawing()
	rl.ClearBackground(backgroundColour)
	rl.BeginMode2D(camera.cam)

	gameMap.Draw()
	player.Draw()

	rl.DrawFPS(int32(screenWidth/2), int32(screenHeight/2))

	rl.EndMode2D()
	rl.EndDrawing()
}

func init() {
	rl.InitWindow(screenWidth, screenHeight, "Playground")
	rl.SetExitKey(rl.KeyQ)
	rl.SetTargetFPS(60)
	rl.InitAudioDevice()

	audio = NewAudio()
	player = NewPlayer()
	camera = NewCamera(player)
	gameMap = NewGameMap()
}

func quit() {
	if err := player.Close(); err != nil {
		panic(err)
	}
	if err := gameMap.Close(); err != nil {
		panic(err)
	}
	if err := audio.Close(); err != nil {
		panic(err)
	}

	defer func() {
		rl.CloseAudioDevice()
		rl.CloseWindow()
	}()
}

func main() {
	for running {
		input()
		update()
		render()
	}
	quit()
}
