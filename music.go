package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Audio struct {
	paused bool
	music  rl.Music
}

func NewAudio() *Audio {
	a := Audio{
		paused: false,
		music:  rl.LoadMusicStream("res/Audio/Loopable.mp3"),
	}
	rl.PlayMusicStream(a.music)

	return &a
}

func (a *Audio) Input() {
	if rl.IsKeyPressed(rl.KeyM) {
		a.paused = !a.paused
	}
}

func (a *Audio) Update() {
	rl.UpdateMusicStream(a.music)
	if a.paused {
		rl.PauseMusicStream(a.music)
	} else {
		rl.ResumeMusicStream(a.music)
	}
}

func (a *Audio) Close() error {
	rl.UnloadMusicStream(a.music)
	return nil
}
