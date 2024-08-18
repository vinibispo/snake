package main

import (
	"fmt"
	constants "snake/internals"
	"snake/internals/colors"
	"snake/internals/models"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var lastUpdateTime float64

func eventTriggered(interval float64) bool {
	currentTime := rl.GetTime()
	if currentTime-lastUpdateTime >= interval {
		lastUpdateTime = currentTime
		return true
	}
	return false
}

func main() {
	var size int32 = constants.Offset*2 + constants.CellSize*constants.CellCount
	rl.InitWindow(size, size, "Retro Snake")

	rl.SetTargetFPS(60)
	game := models.NewGame()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		if eventTriggered(0.2) {
			game.Update()
		}

		game.Move()
		rl.ClearBackground(colors.Green)
		rect := rl.NewRectangle(constants.Offset-5, constants.Offset-5, constants.CellSize*constants.CellCount+10, constants.CellSize*constants.CellCount+10)
		rl.DrawRectangleLinesEx(rect, 5, colors.DarkGreen)
		rl.DrawText("Retro Snake", constants.Offset-5, 20, 40, colors.DarkGreen)

		rl.DrawText(fmt.Sprint(game.Score), constants.Offset-5, constants.Offset+constants.CellSize*constants.CellCount+10, 40, colors.DarkGreen)
		game.Draw()
		rl.EndDrawing()
	}
}
