package main

import (
	"slices"
	"testing"
)

func TestCheckInsertion(t *testing.T) {
	slice1 := []int{10, 20, 30, 40, 20, 50}
	slice2 := []int{20, 40, 60}

	gotBool, gotSlice := CheckInsertion(slice1, slice2)

	wantBool := true
	wantSlice := []int{20, 40}

	if gotBool != wantBool {
		t.Errorf("got bool %v\n want %v\n given %v\n %v", gotBool, wantBool, slice1, slice2)
	}

	if !slices.Equal(gotSlice, wantSlice) {
		t.Errorf("got slice %v\n want %v\n given %v\n %v", gotSlice, wantSlice, slice1, slice2)
	}
}

func TestCheckInsertionNoCommon(t *testing.T) {
	slice1 := []int{1, 2, 3}
	slice2 := []int{4, 5, 6}

	gotBool, gotSlice := CheckInsertion(slice1, slice2)

	wantBool := false

	if gotBool != wantBool {
		t.Errorf("got %v\n want %v", gotBool, wantBool)
	}

	if len(gotSlice) != 0 {
		t.Errorf("got %v\n want empty slice", gotSlice)
	}
}
