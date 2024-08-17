package main

import (
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
	var size int32 = constants.CellSize * constants.CellCount
	rl.InitWindow(size, size, "Retro Snake")

	rl.SetTargetFPS(60)
	game := models.NewGame()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		if eventTriggered(0.2) {
			game.Update()
		}

		game.Snake.Move()
		rl.ClearBackground(colors.Green)
		game.Draw()
		rl.EndDrawing()
	}
}
