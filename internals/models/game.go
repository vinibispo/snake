package models

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	Snake *Snake
	Food  *Food
}

func NewGame() *Game {
	g := &Game{
		Snake: NewSnake(),
	}
	g.Food = NewFood(g.Snake.Body)
	return g
}

func (g *Game) Draw() {
	g.Food.Draw()
	g.Snake.Draw()
}

func (g *Game) Update() {
	g.Snake.Update()
	g.CheckCollisionWithFood()
}

func (g *Game) Move() {
	g.Snake.Move()
}

func (g *Game) CheckCollisionWithFood() {
	if rl.Vector2Equals(g.Snake.Body[0], g.Food.Position) {
		g.Food.Position = g.Food.GenRandomPos(g.Snake.Body)
	}
}
