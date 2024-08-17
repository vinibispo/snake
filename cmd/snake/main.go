package main

import (
	constants "snake/internals"
	"snake/internals/colors"
	"snake/internals/models"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	var size int32 = constants.CellSize * constants.CellCount
	rl.InitWindow(size, size, "Retro Snake")

	rl.SetTargetFPS(60)
	food := models.NewFood()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(colors.Green)
		food.Draw()
		rl.EndDrawing()
	}
}
