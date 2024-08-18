package models

import (
	constants "snake/internals"
	"snake/internals/colors"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Snake struct {
	Body       []rl.Vector2
	Direction  rl.Vector2
	addSegment bool
}

func NewSnake() *Snake {
	s := &Snake{
		Body: []rl.Vector2{
			rl.NewVector2(6, 9),
			rl.NewVector2(5, 9),
			rl.NewVector2(4, 9),
		},
		Direction: rl.NewVector2(1, 0),
	}
	return s
}

func (s *Snake) Draw() {
	for _, part := range s.Body {
		x := part.X
		y := part.Y
		segment := rl.NewRectangle(constants.Offset+x*constants.CellSize, constants.Offset+y*constants.CellSize, constants.CellSize, constants.CellSize)
		rl.DrawRectangleRounded(segment, 0.5, 6, colors.DarkGreen)
	}
}

func (s *Snake) Update() {
	head := s.Body[0]
	newHead := rl.NewVector2(head.X+s.Direction.X, head.Y+s.Direction.Y)
	s.Body = append([]rl.Vector2{newHead}, s.Body...)
	if s.addSegment {
		s.addSegment = false
	} else {
		s.Body = s.Body[:len(s.Body)-1]
	}
}
