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

	var lastFrame Frame
	if g.progress > 0 {
		lastFrame = g.frames[g.progress-1]
		lastFrameWasStrike := lastFrame.firstRollScore == 10
		if lastFrameWasStrike {
			g.score += pins * 2
			return
		}
	}

	if frame.progress == FirstRoll {
		frame.firstRollScore = pins

		if pins == 10 {
			g.frames[g.progress] = frame
			g.score += pins
			g.progress += 1
			return
		}
		lastFrameWasSpare := lastFrame.firstRollScore+lastFrame.secondRollScore == 10
		if lastFrameWasSpare {
			g.score += pins * 2
			return
		}
		frame.progress = SecondRoll
		g.frames[g.progress] = frame
		g.score += pins
	} else {
		frame.secondRollScore = pins
		g.frames[g.progress] = frame
		g.score += pins
		g.progress += 1
	}
}

func (g *Game) Score() int {
	return g.score
}
