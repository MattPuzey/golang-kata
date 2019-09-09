package bowling

type Game struct{}

func New() *Game {
	return &Game{}
}

func (g *Game) Roll(pins int) {
}

func (g *Game) Score() int {
	return 0
}
