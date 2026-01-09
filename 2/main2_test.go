package main

import (
	"slices"
	"testing"
)

func TestSliceExample(t *testing.T) {
	originalSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	got := sliceExample(originalSlice)
	want := []int{2, 4, 6, 8, 10}

	if !slices.Equal(got, want) {
		t.Errorf("sliceExample() got %v, want %v", got, want)
	}
}

func TestAddElements(t *testing.T) {
	originalSlice := []int{1, 2, 3}
	number := 4

	got := addElements(originalSlice, number)
	want := []int{1, 2, 3, 4}

	if !slices.Equal(got, want) {
		t.Errorf("addElements() got %v, want %v", got, want)
	}
}

func TestRemoveElement(t *testing.T) {
	originalSlice := []int{10, 20, 30, 40}
	index := 1 // удаляем 20

	got := removeElement(originalSlice, index)
	want := []int{10, 30, 40}

	if !slices.Equal(got, want) {
		t.Errorf("removeElement() got %v, want %v, index %d", got, want, index)
	}
}

func TestCopySlice(t *testing.T) {
	originalSlice := []int{5, 10, 15}

	got := copySlice(originalSlice)

	if !slices.Equal(got, originalSlice) {
		t.Errorf("copySlice() got %v, want %v", got, originalSlice)
	}

	if len(got) > 0 && &got[0] == &originalSlice[0] {
		t.Errorf("copySlice() didn't return a copy")
	}
}

func TestGenerateSlice(t *testing.T) {
	got := generateSlice()

	if len(got) != 10 {
		t.Errorf("generateSlice() want 10, got %d", len(got))
	}

	if cap(got) != 10 {
		t.Errorf("generateSlice() want 10, got %d", cap(got))
	}
}
