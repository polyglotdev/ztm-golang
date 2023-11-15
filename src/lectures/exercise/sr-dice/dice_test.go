package main

import (
	"testing"
)

func TestRoll(t *testing.T) {
	sides := 6
	for i := 0; i < 100; i++ {
		rolled := roll(sides)
		if rolled < 1 || rolled > sides {
			t.Errorf("Expected roll to be between 1 and %d, got %d", sides, rolled)
		}
	}
}

func TestIsSnakeEyes(t *testing.T) {
	if !isSnakeEyes(2, 2) {
		t.Errorf("Expected isSnakeEyes(2, 2) to return true, got false")
	}
	if isSnakeEyes(2, 3) {
		t.Errorf("Expected isSnakeEyes(2, 3) to return false, got true")
	}
}

func TestIsLuckySeven(t *testing.T) {
	if !isLuckySeven(7) {
		t.Errorf("Expected isLuckySeven(7) to return true, got false")
	}
	if isLuckySeven(6) {
		t.Errorf("Expected isLuckySeven(6) to return false, got true")
	}
}

func TestIsEven(t *testing.T) {
	if !isEven(2) {
		t.Errorf("Expected isEven(2) to return true, got false")
	}
	if isEven(3) {
		t.Errorf("Expected isEven(3) to return false, got true")
	}
}

func TestIsOdd(t *testing.T) {
	if !isOdd(3) {
		t.Errorf("Expected isOdd(3) to return true, got false")
	}
	if isOdd(2) {
		t.Errorf("Expected isOdd(2) to return false, got true")
	}
}

func TestPrintResult(t *testing.T) {
	r := 1
	d := 1
	rolled := 2
	sum := 2
	dice := 2
	printResult(r, d, rolled, sum, dice)
}
