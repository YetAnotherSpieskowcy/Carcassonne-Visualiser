package addons

import rl "github.com/gen2brain/raylib-go/raylib"

type Info struct {
	text string

	position rl.Vector2
}

func NewInfo(text string, pos rl.Vector2) Info {
	return Info{
		text:     text,
		position: pos,
	}
}

func (info Info) Show() {
	rl.DrawText(info.text, int32(info.position.X), int32(info.position.Y), 20, rl.Black)
}
