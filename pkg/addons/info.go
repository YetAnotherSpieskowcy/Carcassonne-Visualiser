package addons

import rl "github.com/gen2brain/raylib-go/raylib"

type Info struct {
	Text string

	Position rl.Vector2
}

func NewInfo(text string, pos rl.Vector2) Info {
	return Info{
		Text:     text,
		Position: pos,
	}
}

func (info Info) DrawInfo() {
	rl.DrawText(info.Text, int32(info.Position.X), int32(info.Position.Y), 20, rl.Black)
}
