package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	screenWidth  = 1_000
	screenHeight = 480
)

var (
	running          bool     = true
	backgroundColour rl.Color = rl.NewColor(147, 211, 196, 255)

	grassSprite  rl.Texture2D
	playerSprite rl.Texture2D

	playerSrc                                     rl.Rectangle
	playerDest                                    rl.Rectangle
	playerMoving                                  bool
	playerDir                                     int
	playerUp, playerDown, playerRight, playerLeft bool
	playerFrame                                   int

	frameCount int

	playerSpeed float32 = 3.0

	musicPaused bool = false
	music       rl.Music

	camera  rl.Camera2D
	camZoom float32 = 1.0

	playAlong bool = true
)

func drawScene() {
	rl.DrawTexture(grassSprite, 100, 50, rl.White)
	rl.DrawTexturePro(playerSprite, playerSrc, playerDest, rl.NewVector2(playerDest.Width, playerDest.Height), 0.0, rl.White)
}

func input() {
	// Movement
	if rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyUp) {
		playerMoving, playerUp, playerDir = true, true, 1
	}
	if rl.IsKeyDown(rl.KeyS) || rl.IsKeyDown(rl.KeyDown) {
		playerMoving, playerDown, playerDir = true, true, 0
	}
	if rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight) {
		playerMoving, playerRight, playerDir = true, true, 3
	}
	if rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft) {
		playerMoving, playerLeft, playerDir = true, true, 2
	}
	// Camera zoom
	if rl.IsKeyDown(rl.KeyKpAdd) || rl.IsKeyDown(rl.KeyZ) {
		if camZoom += 0.1; camZoom > 3.0 {
			camZoom = 3.0
		}
	}
	if rl.IsKeyDown(rl.KeyKpSubtract) || rl.IsKeyDown(rl.KeyX) {
		if camZoom -= 0.1; camZoom < 0.5 {
			camZoom = 0.5
		}
	}
	// Screenshot?
	if rl.IsKeyPressed(rl.KeyO) {
		rl.TakeScreenshot("screenshot.png")
	}
	// Pause music?
	if rl.IsKeyPressed(rl.KeyP) {
		musicPaused = !musicPaused
	}
	// Quit?
	if rl.IsKeyPressed(rl.KeyQ) {
		playAlong = false
	}
}

func update() {
	running = playAlong && !rl.WindowShouldClose()
	playerSrc.X = 0

	if playerMoving {
		if playerUp {
			playerDest.Y -= playerSpeed
		}
		if playerDown {
			playerDest.Y += playerSpeed
		}
		if playerRight {
			playerDest.X += playerSpeed
		}
		if playerLeft {
			playerDest.X -= playerSpeed
		}
		if frameCount%8 == 1 {
			playerFrame++
		}
		playerSrc.X = playerSrc.Width * float32(playerFrame)
	}
	frameCount++
	if playerFrame > 3 {
		playerFrame = 0
	}
	playerSrc.Y = playerSrc.Height * float32(playerDir)

	rl.UpdateMusicStream(music)
	if musicPaused {
		rl.PauseMusicStream(music)
	} else {
		rl.ResumeMusicStream(music)
	}
	camera.Target = rl.NewVector2(playerDest.X-(playerDest.Width/2), playerDest.Y-(playerDest.Height/2))
	camera.Zoom = camZoom
	playerMoving, playerUp, playerDown, playerRight, playerLeft = false, false, false, false, false
}

func render() {
	rl.BeginDrawing()
	rl.ClearBackground(backgroundColour)
	rl.BeginMode2D(camera)

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
	playerSprite = rl.LoadTexture("res/Characters/BasicCharakterSpritesheet.png")

	playerSrc = rl.NewRectangle(0, 0, 48, 48)
	playerDest = rl.NewRectangle(200, 200, 100, 100)

	rl.InitAudioDevice()
	music = rl.LoadMusicStream("res/Audio/Loopable.mp3")
	rl.PlayMusicStream(music)

	camera = rl.NewCamera2D(rl.NewVector2(float32(screenWidth/2), float32(screenHeight/2)),
		rl.NewVector2(playerDest.X-(playerDest.Width/2), playerDest.Y-(playerDest.Height/2)), 0.0, 1.0)
}

func quit() {
	rl.UnloadTexture(grassSprite)
	rl.UnloadTexture(playerSprite)
	rl.UnloadMusicStream(music)
	rl.CloseAudioDevice()
	rl.CloseWindow()
}

func main() {

	for running {
		input()
		update()
		render()
	}

	quit()
}
