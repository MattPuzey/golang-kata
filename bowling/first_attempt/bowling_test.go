package bowling_test

import (
	bowling_game "kata/bowling/first_attempt"
	"testing"
)

func TestGutterGame(t *testing.T) {
	game := bowling_game.New()
	if game == nil {
		t.Fatalf("")
	}
}
