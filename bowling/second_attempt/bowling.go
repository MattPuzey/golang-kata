package bowling

type Game struct {
}

func New() *Game {
	return &Game{}
}

func (g *Game) Bowl(pins int) {}
