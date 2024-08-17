package models

type Game struct {
	Snake *Snake
	Food  *Food
}

func NewGame() *Game {
	g := &Game{
		Snake: NewSnake(),
		Food:  NewFood(),
	}
	return g
}

func (g *Game) Draw() {
	g.Food.Draw()
	g.Snake.Draw()
}

func (g *Game) Update() {
	g.Snake.Update()
}

func (g *Game) Move() {
	g.Snake.Move()
}
