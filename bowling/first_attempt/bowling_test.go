package bowling_test

import (
	game "kata/bowling/first_attempt"
	"testing"
)

func TestGutterGame(t *testing.T) {
	g := game.New()
	if g == nil {
		t.Fatalf("")
	}
}
