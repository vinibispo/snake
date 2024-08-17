package models

import (
	constants "snake/internals"
	"snake/internals/helpers"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Food struct {
	Position rl.Vector2
	Texture  rl.Texture2D
}

func NewFood(snakeBody []rl.Vector2) *Food {
	image := rl.LoadImage("assets/food.png")
	texture := rl.LoadTextureFromImage(image)
	rl.UnloadImage(image)
	f := &Food{
		Position: rl.NewVector2(5, 6),
		Texture:  texture,
	}
	f.Position = f.GenRandomPos(snakeBody)
	return f
}

func (f *Food) GenRandomCell() rl.Vector2 {
	x := rl.GetRandomValue(0, constants.CellCount-1)
	y := rl.GetRandomValue(0, constants.CellCount-1)
	return rl.NewVector2(float32(x), float32(y))
}

func (f *Food) GenRandomPos(snakeBody []rl.Vector2) rl.Vector2 {
	position := f.GenRandomCell()
	for helpers.ElementInSlice(position, snakeBody) {
		position = f.GenRandomCell()
	}
	return position
}

func (f *Food) Draw() {
	rl.DrawTexture(f.Texture, int32(f.Position.X)*constants.CellSize, int32(f.Position.Y)*constants.CellSize, rl.White)
}
