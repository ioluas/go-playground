package main

import (
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type GameMap struct {
	sprites       map[string]rl.Texture2D
	layout        []map[string]string
	src           rl.Rectangle
	dest          rl.Rectangle
	width, height int
}

func NewGameMap() *GameMap {
	m := GameMap{
		sprites: map[string]rl.Texture2D{
			"g": rl.LoadTexture("res/Tilesets/Grass.png"),
			"f": rl.LoadTexture("res/Tilesets/Fences.png"),
			"h": rl.LoadTexture("res/Tilesets/Hills.png"),
			"t": rl.LoadTexture("res/Tilesets/TilledDirt.png"),
			"w": rl.LoadTexture("res/Tilesets/Water.png"),
			"s": rl.LoadTexture("res/Tilesets/WoodenHouse.png"),
		},
		layout: []map[string]string{
			{"t": "w", "v": "2"}, {"t": "w", "v": "1"}, {"t": "w", "v": "2"}, {"t": "w", "v": "1"}, {"t": "w", "v": "2"}, {"t": "w", "v": "1"}, {"t": "w", "v": "2"}, {"t": "w", "v": "1"}, {"t": "w", "v": "2"}, {"t": "w", "v": "1"}, {"t": "w", "v": "2"}, {"t": "w", "v": "1"}, {"t": "w", "v": "2"}, {"t": "w", "v": "1"}, {"t": "w", "v": "2"}, {"t": "w", "v": "1"}, {"t": "w", "v": "2"}, {"t": "w", "v": "1"}, {"t": "w", "v": "2"}, {"t": "w", "v": "1"}, {"t": "w", "v": "2"}, {"t": "w", "v": "1"}, {"t": "w", "v": "2"}, {"t": "w", "v": "1"}, {"t": "w", "v": "2"}, {"t": "w", "v": "1"},
			{"t": "w", "v": "1"}, {"t": "g", "v": "32"}, {"t": "g", "v": "33"}, {"t": "g", "v": "33"}, {"t": "g", "v": "33"}, {"t": "g", "v": "33"}, {"t": "g", "v": "33"}, {"t": "g", "v": "33"}, {"t": "g", "v": "33"}, {"t": "g", "v": "33"}, {"t": "g", "v": "33"}, {"t": "g", "v": "33"}, {"t": "g", "v": "33"}, {"t": "g", "v": "33"}, {"t": "g", "v": "33"}, {"t": "g", "v": "33"}, {"t": "g", "v": "33"}, {"t": "g", "v": "33"}, {"t": "g", "v": "33"}, {"t": "g", "v": "33"}, {"t": "g", "v": "33"}, {"t": "g", "v": "33"}, {"t": "g", "v": "33"}, {"t": "g", "v": "33"}, {"t": "g", "v": "34"}, {"t": "w", "v": "2"},
			{"t": "w", "v": "2"}, {"t": "g", "v": "42"}, {"t": "s", "v": "8"}, {"t": "s", "v": "9"}, {"t": "s", "v": "2"}, {"t": "s", "v": "9"}, {"t": "s", "v": "9"}, {"t": "s", "v": "10"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "44"}, {"t": "w", "v": "1"},
			{"t": "w", "v": "1"}, {"t": "g", "v": "42"}, {"t": "s", "v": "15"}, {"t": "s", "v": "16"}, {"t": "s", "v": "16"}, {"t": "s", "v": "16"}, {"t": "s", "v": "16"}, {"t": "s", "v": "17"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "44"}, {"t": "w", "v": "2"},
			{"t": "w", "v": "2"}, {"t": "g", "v": "42"}, {"t": "s", "v": "15"}, {"t": "s", "v": "16"}, {"t": "s", "v": "16"}, {"t": "s", "v": "16"}, {"t": "s", "v": "16"}, {"t": "s", "v": "17"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "f", "v": "5"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "44"}, {"t": "w", "v": "1"},
			{"t": "w", "v": "2"}, {"t": "g", "v": "42"}, {"t": "s", "v": "15"}, {"t": "s", "v": "16"}, {"t": "s", "v": "16"}, {"t": "s", "v": "16"}, {"t": "s", "v": "16"}, {"t": "s", "v": "17"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "f", "v": "5"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "44"}, {"t": "w", "v": "2"},
			{"t": "w", "v": "2"}, {"t": "g", "v": "42"}, {"t": "s", "v": "22"}, {"t": "s", "v": "23"}, {"t": "s", "v": "11"}, {"t": "s", "v": "23"}, {"t": "s", "v": "23"}, {"t": "s", "v": "24"}, {"t": "f", "v": "3"}, {"t": "f", "v": "3"}, {"t": "f", "v": "3"}, {"t": "f", "v": "3"}, {"t": "f", "v": "4"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "44"}, {"t": "w", "v": "1"},
			{"t": "w", "v": "1"}, {"t": "g", "v": "42"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "44"}, {"t": "w", "v": "2"},
			{"t": "w", "v": "2"}, {"t": "g", "v": "42"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "44"}, {"t": "w", "v": "1"},
			{"t": "w", "v": "1"}, {"t": "g", "v": "42"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "44"}, {"t": "w", "v": "2"},
			{"t": "w", "v": "2"}, {"t": "g", "v": "42"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "45"}, {"t": "w", "v": "53"}, {"t": "w", "v": "53"}, {"t": "w", "v": "53"}, {"t": "g", "v": "46"}, {"t": "g", "v": "44"}, {"t": "w", "v": "1"},
			{"t": "w", "v": "1"}, {"t": "g", "v": "42"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "w", "v": "1"}, {"t": "w", "v": "44"}, {"t": "w", "v": "1"}, {"t": "w", "v": "2"}, {"t": "w", "v": "1"}, {"t": "g", "v": "42"}, {"t": "g", "v": "44"}, {"t": "w", "v": "2"},
			{"t": "w", "v": "2"}, {"t": "g", "v": "42"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "45"}, {"t": "g", "v": "53"}, {"t": "g", "v": "54"}, {"t": "w", "v": "2"}, {"t": "w", "v": "1"}, {"t": "w", "v": "2"}, {"t": "g", "v": "42"}, {"t": "g", "v": "44"}, {"t": "w", "v": "1"},
			{"t": "w", "v": "2"}, {"t": "g", "v": "42"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "1"}, {"t": "g", "v": "44"}, {"t": "w", "v": "1"}, {"t": "w", "v": "2"}, {"t": "w", "v": "1"}, {"t": "w", "v": "2"}, {"t": "w", "v": "1"}, {"t": "g", "v": "42"}, {"t": "g", "v": "44"}, {"t": "w", "v": "1"},
			{"t": "w", "v": "2"}, {"t": "g", "v": "52"}, {"t": "g", "v": "53"}, {"t": "g", "v": "53"}, {"t": "g", "v": "53"}, {"t": "g", "v": "53"}, {"t": "g", "v": "53"}, {"t": "g", "v": "53"}, {"t": "g", "v": "53"}, {"t": "g", "v": "53"}, {"t": "g", "v": "53"}, {"t": "g", "v": "53"}, {"t": "g", "v": "53"}, {"t": "g", "v": "53"}, {"t": "g", "v": "53"}, {"t": "g", "v": "53"}, {"t": "g", "v": "53"}, {"t": "g", "v": "54"}, {"t": "w", "v": "2"}, {"t": "w", "v": "1"}, {"t": "w", "v": "2"}, {"t": "w", "v": "1"}, {"t": "w", "v": "2"}, {"t": "g", "v": "52"}, {"t": "g", "v": "54"}, {"t": "w", "v": "1"},
			{"t": "w", "v": "1"}, {"t": "w", "v": "2"}, {"t": "w", "v": "1"}, {"t": "w", "v": "2"}, {"t": "w", "v": "1"}, {"t": "w", "v": "2"}, {"t": "w", "v": "1"}, {"t": "w", "v": "2"}, {"t": "w", "v": "1"}, {"t": "w", "v": "2"}, {"t": "w", "v": "1"}, {"t": "w", "v": "2"}, {"t": "w", "v": "1"}, {"t": "w", "v": "2"}, {"t": "w", "v": "1"}, {"t": "w", "v": "2"}, {"t": "w", "v": "1"}, {"t": "w", "v": "2"}, {"t": "w", "v": "1"}, {"t": "w", "v": "2"}, {"t": "w", "v": "1"}, {"t": "w", "v": "2"}, {"t": "w", "v": "1"}, {"t": "w", "v": "2"}, {"t": "w", "v": "1"}, {"t": "w", "v": "2"},
		},
		src:    rl.NewRectangle(0, 0, 16, 16),
		dest:   rl.NewRectangle(0, 0, 16, 16),
		width:  26,
		height: 16,
	}

	return &m
}

func (gm *GameMap) Draw() {
	for i := 0; i < len(gm.layout); i++ {
		t := gm.layout[i]["t"]
		sprite := gm.sprites[t]
		v, _ := strconv.Atoi(gm.layout[i]["v"])
		if v != 0 {
			gm.dest.X = gm.dest.Width * float32(i%gm.width)
			gm.dest.Y = gm.dest.Height * float32(i/gm.width)
			if t == "s" || t == "f" {
				rl.DrawTexturePro(gm.sprites["g"], rl.NewRectangle(0, 0, 16, 16), gm.dest,
					rl.NewVector2(gm.dest.Width, gm.dest.Height), 0, rl.White)
			}
			gm.src.X = gm.src.Width * float32((v-1)%int(sprite.Width/int32(gm.src.Width)))
			gm.src.Y = gm.src.Height * float32((v-1)/int(sprite.Width/int32(gm.src.Width)))
			rl.DrawTexturePro(sprite, gm.src, gm.dest, rl.NewVector2(gm.dest.Width, gm.dest.Height), 0, rl.White)
		}
	}
}

func (gm *GameMap) Close() error {
	for _, t := range gm.sprites {
		rl.UnloadTexture(t)
	}
	return nil
}
