package bowling

type Game struct {
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
		progress: 0,
	}
}

func (g *Game) Roll(pins int) {
	progress := g.progress
	frame := g.frames[progress]

	var lastFrame Frame
	var lastFrameWasStrike bool
	if g.progress > 0 {
		lastFrame = g.frames[g.progress-1]
		lastFrameWasStrike = lastFrame.firstRollScore == 10
	}

	if frame.progress == FirstRoll {
		if lastFrameWasStrike {
			frame.firstRollScore += pins * 2
			g.frames[g.progress] = frame
			return
		}

		if pins == 10 {
			frame.firstRollScore = pins
			g.frames[g.progress] = frame
			g.progress += 1
			return
		}
		lastFrameWasSpare := lastFrame.firstRollScore+lastFrame.secondRollScore == 10
		if lastFrameWasSpare {
			frame.firstRollScore += pins * 2
			g.frames[g.progress] = frame
			return
		}
		frame.firstRollScore = pins
		frame.progress = SecondRoll
		g.frames[g.progress] = frame
		return
	}

	if lastFrameWasStrike {
		frame.secondRollScore += pins * 2
		g.frames[g.progress] = frame
		g.progress += 1
		return
	}
	frame.secondRollScore = pins
	g.frames[g.progress] = frame
	g.progress += 1

}

func (g *Game) Score() int {

	var score int
	for _, frame := range g.frames {
		score += frame.firstRollScore + frame.secondRollScore
	}
	return score
}
