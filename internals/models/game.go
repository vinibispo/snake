package models

import (
	constants "snake/internals"
	"snake/internals/helpers"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	Snake     *Snake
	Food      *Food
	running   bool
	Score     int
	eatSound  rl.Sound
	wallSound rl.Sound
}

func NewGame() *Game {
	rl.InitAudioDevice()
	g := &Game{
		Snake:   NewSnake(),
		running: true,
	}
	g.Food = NewFood(g.Snake.Body)
	g.eatSound = rl.LoadSound("assets/eat.mp3")
	g.wallSound = rl.LoadSound("assets/wall.mp3")
	return g
}

func (g *Game) Draw() {
	g.Food.Draw()
	g.Snake.Draw()
}

func (g *Game) Update() {
	if g.running {
		g.Snake.Update()
		g.CheckCollisionWithFood()
		g.CheckCollisionWithEdges()
		g.CheckCollisionWithTail()
	}
}

func (g *Game) Move() {
	s := g.Snake
	if rl.IsKeyPressed(rl.KeyUp) && s.Direction.Y != 1 {
		s.Direction = rl.NewVector2(0, -1)
		g.running = true
	}

	if rl.IsKeyPressed(rl.KeyDown) && s.Direction.Y != -1 {
		s.Direction = rl.NewVector2(0, 1)
		g.running = true
	}

	if rl.IsKeyPressed(rl.KeyLeft) && s.Direction.X != 1 {
		s.Direction = rl.NewVector2(-1, 0)
		g.running = true
	}

	if rl.IsKeyPressed(rl.KeyRight) && s.Direction.X != -1 {
		s.Direction = rl.NewVector2(1, 0)
		g.running = true
	}
}

func (g *Game) CheckCollisionWithFood() {
	if rl.Vector2Equals(g.Snake.Body[0], g.Food.Position) {
		g.Food.Position = g.Food.GenRandomPos(g.Snake.Body)
		g.Snake.addSegment = true
		g.Score++
		rl.PlaySound(g.eatSound)
	}
}

func (g *Game) CheckCollisionWithEdges() {
	if g.Snake.Body[0].X == constants.CellCount || g.Snake.Body[0].X < 0 {
		g.Over()
	}

	if g.Snake.Body[0].Y == constants.CellCount || g.Snake.Body[0].Y < 0 {
		g.Over()
	}
}

func (g *Game) CheckCollisionWithTail() {
	headlessBody := g.Snake.Body[1:]
	if helpers.ElementInSlice(g.Snake.Body[0], headlessBody) {
		g.Over()
	}
}

func (g *Game) Over() {
	g.Snake = NewSnake()
	g.Food = NewFood(g.Snake.Body)
	g.running = false
	g.Score = 0
	rl.PlaySound(g.wallSound)
}
