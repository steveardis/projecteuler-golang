package main

import (
	"fmt"
	"testing"
)

//-----------------------------------------------------------------------------

/*
Should match 3, 5, 6, and 9
*/
func TestP0001AsProvided(t *testing.T) {
	multiplesOf := []int{3, 5}

	result := P0001(10, multiplesOf)
	expected := 23

	if result != expected {
		t.Errorf("result was %d, but expected %d", result, expected)
	}
}

/*
Should match 3, 5, 6, 9, 10, 12, and 15... but do not add 15 to the sum twice (since it is a multiple of both 3 and 5)
*/
func TestP0001WithNumberMatchingMoreThanOneMultiple(t *testing.T) {
	multiplesOf := []int{3, 5}

	result := P0001(16, multiplesOf)
	expected := 60

	if result != expected {
		t.Errorf("result was %d, but expected %d", result, expected)
	}
}

func TestP0001Challenge(t *testing.T) {
	multiplesOf := []int{3, 5}

	fmt.Println("The answer to Problem #1 =", P0001(1000, multiplesOf))
}
