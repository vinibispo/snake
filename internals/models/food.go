package models

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	constants "snake/internals"
)

type Food struct {
	Position rl.Vector2
	Texture  rl.Texture2D
}

func NewFood() *Food {
	image := rl.LoadImage("assets/food.png")
	texture := rl.LoadTextureFromImage(image)
	rl.UnloadImage(image)
	f := &Food{
		Position: rl.NewVector2(5, 6),
		Texture:  texture,
	}
	f.Position = f.Randomize()
	return f
}

func (f *Food) Randomize() rl.Vector2 {
	x := rl.GetRandomValue(0, constants.CellCount-1)
	y := rl.GetRandomValue(0, constants.CellCount-1)
	return rl.NewVector2(float32(x), float32(y))
}

func (f *Food) Draw() {
	rl.DrawTexture(f.Texture, int32(f.Position.X)*constants.CellSize, int32(f.Position.Y)*constants.CellSize, rl.White)
}
