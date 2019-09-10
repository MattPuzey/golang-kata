package bowling

type Game struct {
	score    int
	frames   [10]Frame
	progress int
}

type Frame struct {
	firstRollScore  int
	secondRollScore int
	progress        RollInFrame
}

type RollInFrame int

const (
	FirstRoll RollInFrame = iota
	SecondRoll
)

func New() *Game {
	return &Game{
		score:    0,
		progress: 0,
	}
}

func (g *Game) Roll(pins int) {
	progress := g.progress
	frame := g.frames[progress]
	if frame.progress == FirstRoll {
		frame.firstRollScore = pins
		if g.progress > 0 {
			lastFrame := g.frames[g.progress-1]
			lastFrameWasSpare := lastFrame.firstRollScore+lastFrame.secondRollScore == 10
			if lastFrameWasSpare {
				g.score += pins
			}
		}
		frame.progress = SecondRoll
		g.frames[g.progress] = frame

	} else {
		frame.secondRollScore = pins
		g.frames[g.progress] = frame
		g.progress += 1
	}

	g.score += pins
}

func (g *Game) Score() int {
	return g.score
}
