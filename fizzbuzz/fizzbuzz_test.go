package fizzbuzz_test

import "testing"
import "kata/fizzbuzz"

func TestThereAreAHundredlines(t *testing.T) {

	output := fizzbuzz.Play()

	if len(output) != 100 {
		t.Fatalf("incorrect output")
	}
}

func TestFirstLineIsOne(t *testing.T) {

	output := fizzbuzz.Play()

	if output[0] != "1" {
		t.Fatalf("incorrect output")
	}
}

func TestLastLineIsBuzz(t *testing.T) {

	output := fizzbuzz.Play()
	if output[99] != "buzz" {
		t.Fatalf("incorrect output")
	}
}

func TestThirdLineIsFizz(t *testing.T) {
	output := fizzbuzz.Play()
	if output[2] != "fizz" {
		t.Fatalf("incorrect output")
	}
}

func TestFithLineIsBuzz(t *testing.T) {
	output := fizzbuzz.Play()
	if output[4] != "buzz" {
		t.Fatalf("incorrect output")
	}
}

func TestFifteenthLineIsFizzBuzz(t *testing.T) {
	output := fizzbuzz.Play()
	if output[14] != "fizzbuzz" {
		t.Fatalf("incorrect output")
	}
}
