package main

import (
	"slices"
	"strings"
	"testing"
)

func TestShowType(t *testing.T) {
	result.Reset()

	val := 42
	err := ShowType(val)

	if err != nil {
		t.Fatalf("want nil, got error: %v", err)
	}

	got := strings.TrimSpace(result.String())
	want := "42"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestParseToRuneSlice(t *testing.T) {
	input := "Go"
	got := ParseToRuneSlice(input)
	want := []rune{'G', 'o'}

	if !slices.Equal(got, want) {
		t.Errorf("got %v, want %v, input %q", got, want, input)
	}
}

func TestAddSalt(t *testing.T) {
	input := []rune("Golang")
	// "Gol" + "go-2024" + "ang"

	got := string(AddSalt(input))
	want := "Golgo-2024ang"

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestHash(t *testing.T) {
	input := []byte("test")
	want := "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08"

	got := Hash(input)

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
