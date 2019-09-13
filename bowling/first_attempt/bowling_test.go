package bowling_test

import (
	bowling_game "kata/bowling/first_attempt"
	"strconv"
	"testing"
)

func TestGutterGame(t *testing.T) {
	game := bowling_game.New()

	for i := 1; i <= 20; i++ {
		game.Roll(0)
	}

	score := game.Score()
	if score != 0 {
		t.Fatalf("score %s is incorrect", strconv.Itoa(score))
	}
}

func TestAllOnesScores20(t *testing.T) {
	game := bowling_game.New()

	for i := 1; i <= 20; i++ {
		game.Roll(1)
	}

	score := game.Score()
	if score != 20 {
		t.Fatalf("score %s is incorrect", strconv.Itoa(score))
	}
}

func TestASpareFollowedBy3PinsThenAllMissesScores16(t *testing.T) {
	game := bowling_game.New()

	game.Roll(4)
	game.Roll(6)

	game.Roll(3)

	for i := 1; i <= 17; i++ {
		game.Roll(0)
	}

	score := game.Score()
	if score != 16 {
		t.Fatalf("score %s is incorrect", strconv.Itoa(score))
	}
}

func TestASpareFollowedBy4PinsThenAllMissesScores18(t *testing.T) {
	game := bowling_game.New()

	game.Roll(4)
	game.Roll(6)

	game.Roll(4)

	for i := 1; i <= 17; i++ {
		game.Roll(0)
	}

	score := game.Score()
	if score != 18 {
		t.Fatalf("score %s is incorrect", strconv.Itoa(score))
	}
}

func TestAStrikeFollowedBy3PinsThen2PinsThenAllMissesScores20(t *testing.T) {
	game := bowling_game.New()

	game.Roll(10)

	game.Roll(3)
	game.Roll(2)

	for i := 1; i <= 16; i++ {
		game.Roll(0)
	}

	score := game.Score()
	if score != 20 {
		t.Fatalf("score %s is incorrect", strconv.Itoa(score))
	}
}

func TestAStrikeFollowedBy3PinsThen4PinsThenAllMissesScores24(t *testing.T) {
	game := bowling_game.New()

	game.Roll(10)

	game.Roll(3)
	game.Roll(4)

	for i := 1; i <= 16; i++ {
		game.Roll(0)
	}

	score := game.Score()
	if score != 24 {
		t.Fatalf("score %s is incorrect", strconv.Itoa(score))
	}
}

func TestPerfectGame(t *testing.T) {
	game := bowling_game.New()

	for i := 1; i <= 12; i++ {
		game.Roll(10)
	}

	score := game.Score()
	if score != 300 {
		t.Fatalf("score %s is incorrect", strconv.Itoa(score))
	}
}
