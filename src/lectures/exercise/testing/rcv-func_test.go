// --Summary:
//
//	Copy your rcv-func solution to this directory and write unit tests.
//
// --Requirements:
// * Write unit tests that ensure:
//   - Health & energy can not go above their maximums
//   - Health & energy can not go below 0
//   - If any of your  tests fail, make the necessary corrections
//     in the copy of your rcv-func solution file.
//
// --Notes:
// * Use `go test -v ./exercise/testing` to run these specific tests
package main

import (
	"testing"
)

func TestHealth(t *testing.T) {
	// Test that health can not go above 100
	p := Player{}
	p.health = 100
	p.heal(10)
	if p.health != 100 {
		t.Errorf("Health should not go above 100")
	}

	// Test that health can not go below 0
	p.health = 0
	p.damage(10)
	if p.health != 0 {
		t.Errorf("Health should not go below 0")
	}
}
