package bowling_test

import (
	bowling_game "kata/bowling/second_attempt"
	"testing"
)

func TestGutterGame(t *testing.T) {
	game := bowling_game.New()

	for i := 1; i <= 20; i++ {
		game.Bowl(0)
	}
}
