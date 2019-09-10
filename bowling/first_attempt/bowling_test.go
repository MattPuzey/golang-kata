package bowling_test

import (
	bowling_game "kata/bowling/first_attempt"
	"testing"
)

func TestGutterGame(t *testing.T) {
	game := bowling_game.New()

	for i := 1; i <= 20; i++ {
		game.Roll(0)
	}

	score := game.Score()
	if score != 0 {
		t.Fatalf("score %s is incorrect", string(score))
	}
}

func TestAllOnesScores20(t *testing.T) {
	game := bowling_game.New()

	for i := 1; i <= 20; i++ {
		game.Roll(1)
	}

	score := game.Score()
	if score != 20 {
		t.Fatalf("score %s is incorrect", string(score))
	}
}

func TestASpareFollowedByThreePinsThenAllMissesScores16(t *testing.T) {
	game := bowling_game.New()

	game.Roll(4)
	game.Roll(6)

	game.Roll(3)

	for i := 1; i <= 17; i++ {
		game.Roll(0)
	}

	score := game.Score()
	if score != 16 {
		t.Fatalf("score %s is incorrect", string(score))
	}
}
