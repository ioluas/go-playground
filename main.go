package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	screenWidth  = 1_000
	screenHeight = 480
)

var (
	running     bool = true
	musicPaused bool = false

	backgroundColour rl.Color = rl.NewColor(147, 211, 196, 255)
	grassSprite      rl.Texture2D
	music            rl.Music

	frameCount int

	camera *Camera
	player *Player
)

func drawScene() {
	rl.DrawTexture(grassSprite, 100, 50, rl.White)
	player.Draw()
}

func input() {
	player.Input()
	camera.Input()

	if rl.IsKeyPressed(rl.KeyO) {
		rl.TakeScreenshot("screenshot.png")
	}

	if rl.IsKeyPressed(rl.KeyP) {
		musicPaused = !musicPaused
	}
	// Quit?
	if rl.IsKeyPressed(rl.KeyQ) {
		running = false
	}
}

func update() {
	running = running && !rl.WindowShouldClose()
	frameCount++

	player.Update()
	camera.Update()

	rl.UpdateMusicStream(music)
	if musicPaused {
		rl.PauseMusicStream(music)
	} else {
		rl.ResumeMusicStream(music)
	}
}

func render() {
	rl.BeginDrawing()
	rl.ClearBackground(backgroundColour)
	rl.BeginMode2D(camera.cam)

	drawScene()

	rl.EndMode2D()
	rl.EndDrawing()
}

func init() {
	rl.InitWindow(screenWidth, screenHeight, "Jad's super awesome game")
	rl.SetExitKey(0)
	rl.SetTargetFPS(60)
	rl.SetTraceLog(rl.LogTrace)

	grassSprite = rl.LoadTexture("res/Tilesets/Grass.png")
	player = newPlayer()
	camera = newCamera(*player)

	rl.InitAudioDevice()
	music = rl.LoadMusicStream("res/Audio/Loopable.mp3")
	rl.PlayMusicStream(music)
}

func quit() {
	rl.UnloadMusicStream(music)
	rl.CloseAudioDevice()
	rl.UnloadTexture(grassSprite)
	if err := player.Close(); err != nil {
		panic(err)
	}

	defer rl.CloseWindow()
}

func main() {

	for running {
		input()
		update()
		render()
	}

	quit()
}
