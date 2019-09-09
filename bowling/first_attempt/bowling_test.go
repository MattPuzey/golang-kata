package bowling_test

import (
	bowling_game "kata/bowling/first_attempt"
	"testing"
)

func TestGutterGame(t *testing.T) {
	game := bowling_game.New()
	if game == nil {
		t.Fatalf("nil game created")
	}

	for i := 1; i <= 20; i++ {
		game.Roll(0)
	}

	score := game.Score()
	if score != 0 {
		t.Fatalf("score %s is incorrect", string(score))
	}
}
