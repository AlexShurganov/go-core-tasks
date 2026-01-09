package main

import (
	"slices"
	"testing"
)

func TestFindCommon(t *testing.T) {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}

	got := FindCommon(slice1, slice2)
	want := []string{"banana", "date"}

	if !slices.Equal(got, want) {
		t.Errorf("got %v\n want %v\n given %v\n%v", got, want, slice1, slice2)
	}
}
