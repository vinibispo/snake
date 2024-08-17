package models

import (
	constants "snake/internals"
	"snake/internals/colors"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Snake struct {
	Body []rl.Vector2
}

func NewSnake() *Snake {
	s := &Snake{
		Body: []rl.Vector2{
			rl.NewVector2(6, 9),
			rl.NewVector2(5, 9),
			rl.NewVector2(4, 9),
		},
	}
	return s
}

func (s *Snake) Draw() {
	for _, part := range s.Body {
		x := part.X
		y := part.Y
		rect := rl.NewRectangle(x*constants.CellSize, y*constants.CellSize, constants.CellSize, constants.CellSize)
		rl.DrawRectangleRounded(rect, 0.5, 6, colors.DarkGreen)
	}
}
