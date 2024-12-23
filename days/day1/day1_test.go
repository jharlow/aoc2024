package day1

import (
	"os"
	"testing"
)

const SampleInput = `3   4
4   3
2   5
1   3
3   9
3   3`

func TestDay1Part1Sample(t *testing.T) {
	// Act
	result := Day1Part1(SampleInput)

	// Assert
	if result != 11 {
		t.Fatalf(`Day1Part1(sampleInput) = %d, want match for %d`, result, 11)
	}
}

func TestDay1Part1Real(t *testing.T) {
	// Assemble
	content, err := os.ReadFile("input.txt")
	if err != nil {
		t.Fatal(err)
	}
	realInput := string(content)

	// Act
	result := Day1Part1(realInput)

	// Assert
	if result != 2375403 {
		t.Fatalf(`Day1Part1(realInput) = %d, want match for %d`, result, 2375403)
	}
}

func TestDay1Part2Sample(t *testing.T) {
	// Act
	result := Day1Part2(SampleInput)

	// Assert
	if result != 31 {
		t.Fatalf(`Day1Part2(sampleInput) = %d, want match for %d`, result, 31)
	}
}

func TestDay1Part2Real(t *testing.T) {
	// Assemble
	content, err := os.ReadFile("input.txt")
	if err != nil {
		t.Fatal(err)
	}
	realInput := string(content)

	// Act
	result := Day1Part2(realInput)

	// Assert
	if result != 23082277 {
		t.Fatalf(`Day1Part2(realInput) = %d, want match for %d`, result, 23082277)
	}
}
