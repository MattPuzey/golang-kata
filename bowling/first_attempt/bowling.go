package bowling

type Game struct {
	score         int
	frames        [10]Frame
	gameProgress  int
	frameProgress int
}

type Frame struct {
	firstRoll  int
	secondRoll int
}

func New() *Game {
	return &Game{
		score:         0,
		gameProgress:  0,
		frameProgress: 0,
	}
}

func (g *Game) Roll(pins int) {
	progress := g.gameProgress
	frame := g.frames[progress]
	if g.frameProgress == 0 {
		frame.firstRoll = pins
		if g.gameProgress > 0 {
			lastFrame := g.frames[g.gameProgress-1]
			lastFrameWasSpare := lastFrame.firstRoll+lastFrame.secondRoll == 10
			if lastFrameWasSpare {
				g.score += pins
			}
		}
		g.frameProgress += 1
		g.frames[g.gameProgress] = frame

	} else {
		frame.secondRoll = pins
		g.frames[g.gameProgress] = frame
		g.gameProgress += 1
		g.frameProgress = 0
	}

	g.score += pins
}

func (g *Game) Score() int {
	return g.score
}
