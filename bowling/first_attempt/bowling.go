package bowling

type Game struct {
	score int
}

func New() *Game {
	return &Game{}
}

func (g *Game) Roll(pins int) {
	if pins != 0 {
		g.score += 1
	}
}

func (g *Game) Score() int {
	return g.score
}
